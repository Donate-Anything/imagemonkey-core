
--Image table
CREATE INDEX image_image_provider_index ON image (image_provider_id);
CREATE INDEX image_key_index ON image (key);
CREATE INDEX image_hash_index ON image (hash);
CREATE INDEX image_unlocked_index ON image (unlocked);

--Image Validation table
CREATE INDEX image_validation_image_id_index ON image_validation (image_id);
CREATE INDEX image_validation_label_id_index ON image_validation (label_id);


-- Label table
CREATE INDEX label_name_index ON label (name);
CREATE INDEX label_parent_id_index ON label(parent_id);

-- annotation_data table
CREATE INDEX annotation_data_image_annotation_id_idx ON annotation_data(image_annotation_id);