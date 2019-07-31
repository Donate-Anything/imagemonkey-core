package main

import (
	"database/sql"
	"flag"
	"fmt"
	commons "github.com/bbernhard/imagemonkey-core/commons"
	datastructures "github.com/bbernhard/imagemonkey-core/datastructures"
	"github.com/getsentry/raven-go"
	"github.com/gofrs/uuid"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"time"
)

var db *sql.DB

func setTrendingLabelBotTaskState(status string, branchName string, jobUrl string, id int64) error {
	var queryValues []interface{}
	queryValues = append(queryValues, status)
	queryValues = append(queryValues, branchName)
	queryValues = append(queryValues, id)
	jobUrlStr := ""
	if jobUrl != "" {
		jobUrlStr = ", job_url = $4"
		queryValues = append(queryValues, jobUrl)
	}

	q := fmt.Sprintf(`UPDATE trending_label_bot_task 
					  	SET state = $1, branch_name = $2%s
						WHERE id = $3`, jobUrlStr)
	_, err := db.Exec(q, queryValues...)
	return err
}

func resetTrendingLabelBotTaskState(id int64) error {
	_, err := db.Exec(`UPDATE trending_label_bot_task
						SET state = 'accepted', branch_name = null, job_url = null,
						try = try + 1
						WHERE id = $1`, id)
	return err
}

func getTrendingLabels() ([]datastructures.TrendingLabelBotTask, error) {
	trendingLabels := []datastructures.TrendingLabelBotTask{}

	rows, err := db.Query(`SELECT s.name, b.id, b.state, COALESCE(b.branch_name, ''), label_type,
						   COALESCE(b.plural) as plural, COALESCE(b.description, ''), COALESCE(b.rename_to, '')
					  	   FROM trending_label_suggestion t
					  	   JOIN trending_label_bot_task b ON b.trending_label_suggestion_id = t.id 
					  	   JOIN label_suggestion s ON s.id = t.label_suggestion_id
						   WHERE t.closed = false AND 
						   (b.state = 'accepted' OR b.state='pending' OR b.state='building' 
						   	OR b.state='build-success' OR b.state='retry')`)
	if err != nil {
		return trendingLabels, err
	}

	defer rows.Close()

	for rows.Next() {
		var trendingLabel datastructures.TrendingLabelBotTask
		err = rows.Scan(&trendingLabel.Name, &trendingLabel.BotTaskId, &trendingLabel.State, 
							&trendingLabel.BranchName, &trendingLabel.LabelType, &trendingLabel.Plural,
							&trendingLabel.Description, &trendingLabel.RenameTo)
		if err != nil {
			return trendingLabels, err
		}

		trendingLabels = append(trendingLabels, trendingLabel)
	}

	return trendingLabels, nil
}

func main() {
	labelsRepositoryName := flag.String("labels_repository_name", "imagemonkey-trending-labels-test", "Label Repository Name")
	labelsRepositoryOwner := flag.String("labels_repository_owner", "bbernhard", "Label Repository Owner")
	metalabelsPath := flag.String("metalabels", "../wordlists/en/metalabels.jsonnet", "Path to metalabels")
	labelsPath := flag.String("labels", "../wordlists/en/labels.jsonnet", "Path to labels")
	gitCheckoutDir := flag.String("git_checkout_dir", "/tmp/labelrepository", "Git checkout directory")

	flag.Parse()

	log.SetLevel(log.DebugLevel)
	log.Info("Starting ImageMonkey Bot")

	imageMonkeyBotGithubApiToken := commons.MustGetEnv("IMAGEMONKEY_BOT_GITHUB_API_TOKEN")
	travisCiApiToken := commons.MustGetEnv("IMAGEMONKEY_TRAVIS_CI_TOKEN")

	log.Debug("Reading Metalabels")
	metaLabels := commons.NewMetaLabels(*metalabelsPath)
	err := metaLabels.Load()
	if err != nil {
		log.Fatal("Couldn't read metalabel map...terminating! ", err.Error())
	}

	log.Debug("Reading Labels")
	labels := commons.NewLabelRepository(*labelsPath)
	err = labels.Load()
	if err != nil {
		log.Fatal("Couldn't read label map...terminating!", err.Error())
	}

	//open database and make sure that we can ping it
	imageMonkeyDbConnectionString := commons.MustGetEnv("IMAGEMONKEY_DB_CONNECTION_STRING")
	db, err = sql.Open("postgres", imageMonkeyDbConnectionString)
	if err != nil {
		log.Fatal("Couldn't open database: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Couldn't ping database: ", err.Error())
	}
	defer db.Close()

	labelsRepository := commons.NewLabelsRepository(*labelsRepositoryOwner, *labelsRepositoryName, *gitCheckoutDir)
	labelsRepository.SetToken(imageMonkeyBotGithubApiToken)

	travisCiApi := commons.NewTravisCiApi("bbernhard", "imagemonkey-trending-labels-test")
	travisCiApi.SetToken(travisCiApiToken)

	firstIteration := true
	for {

		if !firstIteration {
			time.Sleep(1 * time.Second)
		} else {
			firstIteration = false
		}

		log.Debug("Fetch trending labels")

		trendingLabels, err := getTrendingLabels()
		if err != nil {
			log.Error("Couldn't get trending labels: ", err.Error())
			raven.CaptureError(err, nil)
		}
		for _, trendingLabel := range trendingLabels {
			labelsRepository.RemoveLocal()

			if trendingLabel.State == "accepted" {
				log.Info("Got new trending label ", trendingLabel.Name)

				if metaLabels.Contains(trendingLabel.RenameTo) || labels.Contains(trendingLabel.RenameTo, "") {
					err = setTrendingLabelBotTaskState("already exists", trendingLabel.BranchName, "", trendingLabel.BotTaskId)
					if err != nil {
						log.Error("Couldn't set trending label bot task state to 'already exists': ", err.Error())
						raven.CaptureError(err, nil)
					}
					continue //trendinglabel already exists
				}

				err = labelsRepository.Clone()
				if err != nil {
					log.Error(err.Error())
					raven.CaptureError(err, nil)
					continue
				}

				u, err := uuid.NewV4()
				if err != nil {
					log.Error("Couldn't create UUID: ", err.Error())
					raven.CaptureError(err, nil)
					continue
				}
				trendingLabel.Name = u.String()

				branchName, err := labelsRepository.AddLabelAndPushToRepo(trendingLabel)
				if err != nil {
					log.Error("Couldn't add label: ", err.Error())
					raven.CaptureError(err, nil)
					continue
				}

				err = setTrendingLabelBotTaskState("pending", branchName, "", trendingLabel.BotTaskId)
				if err != nil {
					log.Error("Couldn't set trending label bot task state to pending: ", err.Error())
					raven.CaptureError(err, nil)
					continue
				}
				trendingLabel.State = "pending"
				trendingLabel.BranchName = branchName
			}
			if trendingLabel.State == "pending" {
				err = travisCiApi.StartBuild(trendingLabel.BranchName)
				if err != nil {
					log.Error("Couldn't start travis build: ", err.Error())
					raven.CaptureError(err, nil)
					continue
				}

				err = setTrendingLabelBotTaskState("building", trendingLabel.BranchName, "", trendingLabel.BotTaskId)
				if err != nil {
					log.Error("Couldn't set trending label bot task state to building: ", err.Error())
					raven.CaptureError(err, nil)
					continue
				}
			}
			if trendingLabel.State == "building" {
				travisCiBuildInfo, err := travisCiApi.GetBuildInfo(trendingLabel.BranchName)
				if err != nil {
					log.Error("Couldn't query build info for branch ", trendingLabel.BranchName, ": ", err.Error())
					raven.CaptureError(err, nil)
					continue
				}
				if travisCiBuildInfo.LastBuild.State == "created" {
					err = setTrendingLabelBotTaskState("building", trendingLabel.BranchName, travisCiBuildInfo.JobUrl, trendingLabel.BotTaskId)
					if err != nil {
						log.Error("Couldn't set trending label bot task state to buildings: ", err.Error())
						raven.CaptureError(err, nil)
						continue
					}
				} else if travisCiBuildInfo.LastBuild.State == "started" {
					err = setTrendingLabelBotTaskState("building", trendingLabel.BranchName, travisCiBuildInfo.JobUrl, trendingLabel.BotTaskId)
					if err != nil {
						log.Error("Couldn't set trending label bot task state to building: ", err.Error())
						raven.CaptureError(err, nil)
						continue
					}
				} else if travisCiBuildInfo.LastBuild.State == "passed" {
					err = setTrendingLabelBotTaskState("build-success", trendingLabel.BranchName, travisCiBuildInfo.JobUrl, trendingLabel.BotTaskId)
					if err != nil {
						log.Error("Couldn't set trending label bot task state to build-success: ", err.Error())
						raven.CaptureError(err, nil)
						continue
					}
				} else if travisCiBuildInfo.LastBuild.State == "failed" {
					err = setTrendingLabelBotTaskState("build-failed", trendingLabel.BranchName, travisCiBuildInfo.JobUrl, trendingLabel.BotTaskId)
					if err != nil {
						log.Error("Couldn't set trending label bot task state to build-failed: ", err.Error())
						raven.CaptureError(err, nil)
						continue
					}
				} else if travisCiBuildInfo.LastBuild.State == "canceled" {
					err = setTrendingLabelBotTaskState("build-canceled", trendingLabel.BranchName, travisCiBuildInfo.JobUrl, trendingLabel.BotTaskId)
					if err != nil {
						log.Error("Couldn't set trending label bot task state to build-canceled: ", err.Error())
						raven.CaptureError(err, nil)
						continue
					}
				}
			}
			if trendingLabel.State == "build-success" {
				err = labelsRepository.MergeRemoteBranchIntoMaster(trendingLabel.BranchName)
				if err != nil {
					log.Error("Couldn't merge remote branch ", trendingLabel.BranchName, " into master: ", err.Error())
					raven.CaptureError(err, nil)
					continue
				}
				err = setTrendingLabelBotTaskState("merged", trendingLabel.BranchName, "", trendingLabel.BotTaskId)
				if err != nil {
					log.Error("Couldn't set trending label bot task state to merged: ", err.Error())
					raven.CaptureError(err, nil)
					continue
				}
			}
			if trendingLabel.State == "retry" {
				err = labelsRepository.RemoveRemoteBranch(trendingLabel.BranchName)
				if err != nil {
					log.Error("Couldn't remove branch ", trendingLabel.BranchName, ": ", err.Error())
					raven.CaptureError(err, nil)
					continue
				}
				err = resetTrendingLabelBotTaskState(trendingLabel.BotTaskId)
				if err != nil {
					log.Error("Couldn't reset trending label bot task state: ", err.Error())
					raven.CaptureError(err, nil)
					continue
				}
			}
		}
	}
}
