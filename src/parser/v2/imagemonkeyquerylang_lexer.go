// Code generated from ../grammar/ImagemonkeyQueryLang.g4 by ANTLR 4.7.1. DO NOT EDIT.

package imagemonkeyquerylang

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 23, 321,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 5, 19, 157, 10, 19, 3, 20, 6, 20, 160, 10, 20, 13, 20,
	14, 20, 161, 3, 20, 3, 20, 6, 20, 166, 10, 20, 13, 20, 14, 20, 167, 5,
	20, 170, 10, 20, 3, 20, 7, 20, 173, 10, 20, 12, 20, 14, 20, 176, 11, 20,
	3, 20, 3, 20, 7, 20, 180, 10, 20, 12, 20, 14, 20, 183, 11, 20, 3, 20, 3,
	20, 6, 20, 187, 10, 20, 13, 20, 14, 20, 188, 3, 20, 3, 20, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 6, 21, 199, 10, 21, 13, 21, 14, 21, 200, 3,
	21, 3, 21, 3, 21, 3, 22, 3, 22, 6, 22, 208, 10, 22, 13, 22, 14, 22, 209,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 6, 22, 224, 10, 22, 13, 22, 14, 22, 225, 3, 22, 3, 22, 3, 23,
	3, 23, 6, 23, 232, 10, 23, 13, 23, 14, 23, 233, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 6, 23, 248,
	10, 23, 13, 23, 14, 23, 249, 3, 23, 3, 23, 3, 24, 3, 24, 6, 24, 256, 10,
	24, 13, 24, 14, 24, 257, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 7, 25, 278, 10, 25, 12, 25, 14, 25, 281, 11, 25, 3, 25, 3, 25, 5, 25,
	285, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 6, 27, 301, 10, 27, 13, 27, 14,
	27, 302, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32,
	3, 32, 3, 33, 6, 33, 316, 10, 33, 13, 33, 14, 33, 317, 3, 33, 3, 33, 2,
	2, 34, 3, 2, 5, 2, 7, 2, 9, 2, 11, 2, 13, 2, 15, 2, 17, 2, 19, 2, 21, 2,
	23, 2, 25, 3, 27, 4, 29, 5, 31, 6, 33, 7, 35, 8, 37, 9, 39, 10, 41, 11,
	43, 12, 45, 13, 47, 14, 49, 15, 51, 16, 53, 17, 55, 18, 57, 19, 59, 20,
	61, 21, 63, 22, 65, 23, 3, 2, 19, 3, 2, 99, 124, 3, 2, 67, 92, 4, 2, 67,
	92, 99, 124, 5, 2, 34, 34, 67, 92, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124,
	4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103, 4, 2, 85, 85, 117, 117,
	4, 2, 69, 69, 101, 101, 4, 2, 67, 67, 99, 99, 4, 2, 62, 62, 64, 64, 4,
	2, 81, 81, 113, 113, 4, 2, 84, 84, 116, 116, 4, 2, 68, 68, 100, 100, 4,
	2, 91, 91, 123, 123, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 332,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2,
	2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2,
	2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3,
	2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55,
	3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2,
	63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 3, 67, 3, 2, 2, 2, 5, 69, 3, 2, 2, 2,
	7, 71, 3, 2, 2, 2, 9, 73, 3, 2, 2, 2, 11, 75, 3, 2, 2, 2, 13, 80, 3, 2,
	2, 2, 15, 82, 3, 2, 2, 2, 17, 84, 3, 2, 2, 2, 19, 86, 3, 2, 2, 2, 21, 89,
	3, 2, 2, 2, 23, 94, 3, 2, 2, 2, 25, 98, 3, 2, 2, 2, 27, 100, 3, 2, 2, 2,
	29, 120, 3, 2, 2, 2, 31, 132, 3, 2, 2, 2, 33, 145, 3, 2, 2, 2, 35, 147,
	3, 2, 2, 2, 37, 156, 3, 2, 2, 2, 39, 159, 3, 2, 2, 2, 41, 192, 3, 2, 2,
	2, 43, 205, 3, 2, 2, 2, 45, 229, 3, 2, 2, 2, 47, 253, 3, 2, 2, 2, 49, 284,
	3, 2, 2, 2, 51, 286, 3, 2, 2, 2, 53, 300, 3, 2, 2, 2, 55, 304, 3, 2, 2,
	2, 57, 306, 3, 2, 2, 2, 59, 308, 3, 2, 2, 2, 61, 310, 3, 2, 2, 2, 63, 312,
	3, 2, 2, 2, 65, 315, 3, 2, 2, 2, 67, 68, 9, 2, 2, 2, 68, 4, 3, 2, 2, 2,
	69, 70, 9, 3, 2, 2, 70, 6, 3, 2, 2, 2, 71, 72, 9, 4, 2, 2, 72, 8, 3, 2,
	2, 2, 73, 74, 9, 5, 2, 2, 74, 10, 3, 2, 2, 2, 75, 76, 9, 6, 2, 2, 76, 77,
	9, 6, 2, 2, 77, 78, 9, 6, 2, 2, 78, 79, 9, 6, 2, 2, 79, 12, 3, 2, 2, 2,
	80, 81, 7, 34, 2, 2, 81, 14, 3, 2, 2, 2, 82, 83, 7, 97, 2, 2, 83, 16, 3,
	2, 2, 2, 84, 85, 7, 49, 2, 2, 85, 18, 3, 2, 2, 2, 86, 87, 7, 47, 2, 2,
	87, 88, 7, 64, 2, 2, 88, 20, 3, 2, 2, 2, 89, 90, 9, 7, 2, 2, 90, 91, 9,
	8, 2, 2, 91, 92, 9, 9, 2, 2, 92, 93, 9, 10, 2, 2, 93, 22, 3, 2, 2, 2, 94,
	95, 9, 11, 2, 2, 95, 96, 9, 9, 2, 2, 96, 97, 9, 10, 2, 2, 97, 24, 3, 2,
	2, 2, 98, 99, 7, 35, 2, 2, 99, 26, 3, 2, 2, 2, 100, 101, 7, 99, 2, 2, 101,
	102, 7, 112, 2, 2, 102, 103, 7, 112, 2, 2, 103, 104, 7, 113, 2, 2, 104,
	105, 7, 118, 2, 2, 105, 106, 7, 99, 2, 2, 106, 107, 7, 118, 2, 2, 107,
	108, 7, 107, 2, 2, 108, 109, 7, 113, 2, 2, 109, 110, 7, 112, 2, 2, 110,
	111, 7, 48, 2, 2, 111, 112, 7, 101, 2, 2, 112, 113, 7, 113, 2, 2, 113,
	114, 7, 120, 2, 2, 114, 115, 7, 103, 2, 2, 115, 116, 7, 116, 2, 2, 116,
	117, 7, 99, 2, 2, 117, 118, 7, 105, 2, 2, 118, 119, 7, 103, 2, 2, 119,
	28, 3, 2, 2, 2, 120, 121, 7, 107, 2, 2, 121, 122, 7, 111, 2, 2, 122, 123,
	7, 99, 2, 2, 123, 124, 7, 105, 2, 2, 124, 125, 7, 103, 2, 2, 125, 126,
	7, 48, 2, 2, 126, 127, 7, 121, 2, 2, 127, 128, 7, 107, 2, 2, 128, 129,
	7, 102, 2, 2, 129, 130, 7, 118, 2, 2, 130, 131, 7, 106, 2, 2, 131, 30,
	3, 2, 2, 2, 132, 133, 7, 107, 2, 2, 133, 134, 7, 111, 2, 2, 134, 135, 7,
	99, 2, 2, 135, 136, 7, 105, 2, 2, 136, 137, 7, 103, 2, 2, 137, 138, 7,
	48, 2, 2, 138, 139, 7, 106, 2, 2, 139, 140, 7, 103, 2, 2, 140, 141, 7,
	107, 2, 2, 141, 142, 7, 105, 2, 2, 142, 143, 7, 106, 2, 2, 143, 144, 7,
	118, 2, 2, 144, 32, 3, 2, 2, 2, 145, 146, 7, 39, 2, 2, 146, 34, 3, 2, 2,
	2, 147, 148, 7, 114, 2, 2, 148, 149, 7, 122, 2, 2, 149, 36, 3, 2, 2, 2,
	150, 157, 9, 12, 2, 2, 151, 152, 7, 64, 2, 2, 152, 157, 7, 63, 2, 2, 153,
	157, 7, 63, 2, 2, 154, 155, 7, 62, 2, 2, 155, 157, 7, 63, 2, 2, 156, 150,
	3, 2, 2, 2, 156, 151, 3, 2, 2, 2, 156, 153, 3, 2, 2, 2, 156, 154, 3, 2,
	2, 2, 157, 38, 3, 2, 2, 2, 158, 160, 5, 7, 4, 2, 159, 158, 3, 2, 2, 2,
	160, 161, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162,
	169, 3, 2, 2, 2, 163, 165, 7, 48, 2, 2, 164, 166, 5, 7, 4, 2, 165, 164,
	3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167, 168, 3, 2,
	2, 2, 168, 170, 3, 2, 2, 2, 169, 163, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2,
	170, 174, 3, 2, 2, 2, 171, 173, 5, 13, 7, 2, 172, 171, 3, 2, 2, 2, 173,
	176, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 177,
	3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 177, 181, 7, 63, 2, 2, 178, 180, 5, 13,
	7, 2, 179, 178, 3, 2, 2, 2, 180, 183, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2,
	181, 182, 3, 2, 2, 2, 182, 184, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 184,
	186, 7, 41, 2, 2, 185, 187, 5, 9, 5, 2, 186, 185, 3, 2, 2, 2, 187, 188,
	3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 190, 3, 2,
	2, 2, 190, 191, 7, 41, 2, 2, 191, 40, 3, 2, 2, 2, 192, 193, 9, 13, 2, 2,
	193, 194, 9, 14, 2, 2, 194, 195, 9, 7, 2, 2, 195, 196, 9, 8, 2, 2, 196,
	198, 9, 14, 2, 2, 197, 199, 5, 13, 7, 2, 198, 197, 3, 2, 2, 2, 199, 200,
	3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201, 202, 3, 2,
	2, 2, 202, 203, 9, 15, 2, 2, 203, 204, 9, 16, 2, 2, 204, 42, 3, 2, 2, 2,
	205, 207, 5, 41, 21, 2, 206, 208, 5, 13, 7, 2, 207, 206, 3, 2, 2, 2, 208,
	209, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 211,
	3, 2, 2, 2, 211, 212, 7, 120, 2, 2, 212, 213, 7, 99, 2, 2, 213, 214, 7,
	110, 2, 2, 214, 215, 7, 107, 2, 2, 215, 216, 7, 102, 2, 2, 216, 217, 7,
	99, 2, 2, 217, 218, 7, 118, 2, 2, 218, 219, 7, 107, 2, 2, 219, 220, 7,
	113, 2, 2, 220, 221, 7, 112, 2, 2, 221, 223, 3, 2, 2, 2, 222, 224, 5, 13,
	7, 2, 223, 222, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2,
	225, 226, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 228, 5, 21, 11, 2, 228,
	44, 3, 2, 2, 2, 229, 231, 5, 41, 21, 2, 230, 232, 5, 13, 7, 2, 231, 230,
	3, 2, 2, 2, 232, 233, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 233, 234, 3, 2,
	2, 2, 234, 235, 3, 2, 2, 2, 235, 236, 7, 120, 2, 2, 236, 237, 7, 99, 2,
	2, 237, 238, 7, 110, 2, 2, 238, 239, 7, 107, 2, 2, 239, 240, 7, 102, 2,
	2, 240, 241, 7, 99, 2, 2, 241, 242, 7, 118, 2, 2, 242, 243, 7, 107, 2,
	2, 243, 244, 7, 113, 2, 2, 244, 245, 7, 112, 2, 2, 245, 247, 3, 2, 2, 2,
	246, 248, 5, 13, 7, 2, 247, 246, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249,
	247, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 252,
	5, 23, 12, 2, 252, 46, 3, 2, 2, 2, 253, 255, 5, 41, 21, 2, 254, 256, 5,
	13, 7, 2, 255, 254, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 255, 3, 2, 2,
	2, 257, 258, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 260, 7, 120, 2, 2,
	260, 261, 7, 99, 2, 2, 261, 262, 7, 110, 2, 2, 262, 263, 7, 107, 2, 2,
	263, 264, 7, 102, 2, 2, 264, 265, 7, 99, 2, 2, 265, 266, 7, 118, 2, 2,
	266, 267, 7, 107, 2, 2, 267, 268, 7, 113, 2, 2, 268, 269, 7, 112, 2, 2,
	269, 48, 3, 2, 2, 2, 270, 285, 5, 7, 4, 2, 271, 279, 5, 7, 4, 2, 272, 278,
	5, 13, 7, 2, 273, 278, 5, 7, 4, 2, 274, 278, 5, 15, 8, 2, 275, 278, 5,
	17, 9, 2, 276, 278, 5, 19, 10, 2, 277, 272, 3, 2, 2, 2, 277, 273, 3, 2,
	2, 2, 277, 274, 3, 2, 2, 2, 277, 275, 3, 2, 2, 2, 277, 276, 3, 2, 2, 2,
	278, 281, 3, 2, 2, 2, 279, 277, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280,
	282, 3, 2, 2, 2, 281, 279, 3, 2, 2, 2, 282, 283, 5, 7, 4, 2, 283, 285,
	3, 2, 2, 2, 284, 270, 3, 2, 2, 2, 284, 271, 3, 2, 2, 2, 285, 50, 3, 2,
	2, 2, 286, 287, 5, 11, 6, 2, 287, 288, 5, 11, 6, 2, 288, 289, 7, 47, 2,
	2, 289, 290, 5, 11, 6, 2, 290, 291, 7, 47, 2, 2, 291, 292, 5, 11, 6, 2,
	292, 293, 7, 47, 2, 2, 293, 294, 5, 11, 6, 2, 294, 295, 7, 47, 2, 2, 295,
	296, 5, 11, 6, 2, 296, 297, 5, 11, 6, 2, 297, 298, 5, 11, 6, 2, 298, 52,
	3, 2, 2, 2, 299, 301, 9, 17, 2, 2, 300, 299, 3, 2, 2, 2, 301, 302, 3, 2,
	2, 2, 302, 300, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 54, 3, 2, 2, 2,
	304, 305, 7, 40, 2, 2, 305, 56, 3, 2, 2, 2, 306, 307, 7, 126, 2, 2, 307,
	58, 3, 2, 2, 2, 308, 309, 7, 128, 2, 2, 309, 60, 3, 2, 2, 2, 310, 311,
	7, 42, 2, 2, 311, 62, 3, 2, 2, 2, 312, 313, 7, 43, 2, 2, 313, 64, 3, 2,
	2, 2, 314, 316, 9, 18, 2, 2, 315, 314, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2,
	317, 315, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319,
	320, 8, 33, 2, 2, 320, 66, 3, 2, 2, 2, 21, 2, 156, 161, 167, 169, 174,
	181, 188, 200, 209, 225, 233, 249, 257, 277, 279, 284, 302, 317, 3, 8,
	2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'!'", "'annotation.coverage'", "'image.width'", "'image.height'",
	"'%'", "'px'", "", "", "", "", "", "", "", "", "", "'&'", "'|'", "'~'",
	"'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "SEP", "ANNOTATION_COVERAGE_PREFIX", "IMAGE_WIDTH_PREFIX", "IMAGE_HEIGHT_PREFIX",
	"PERCENT", "PIXEL", "OPERATOR", "ASSIGNMENT", "ORDER_BY", "ORDER_BY_VALIDATION_DESC",
	"ORDER_BY_VALIDATION_ASC", "ORDER_BY_VALIDATION", "LABEL", "UUID", "VAL",
	"AND", "OR", "NOT", "LPAR", "RPAR", "SKIPPED_TOKENS",
}

var lexerRuleNames = []string{
	"LOWERCASE", "UPPERCASE", "UPPERLOWERCASE", "UPPERLOWERCASEWS", "UUIDBLOCK",
	"WS", "UNDERSCORE", "SLASH", "GRAPHARROW", "DESC", "ASC", "SEP", "ANNOTATION_COVERAGE_PREFIX",
	"IMAGE_WIDTH_PREFIX", "IMAGE_HEIGHT_PREFIX", "PERCENT", "PIXEL", "OPERATOR",
	"ASSIGNMENT", "ORDER_BY", "ORDER_BY_VALIDATION_DESC", "ORDER_BY_VALIDATION_ASC",
	"ORDER_BY_VALIDATION", "LABEL", "UUID", "VAL", "AND", "OR", "NOT", "LPAR",
	"RPAR", "SKIPPED_TOKENS",
}

type ImagemonkeyQueryLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewImagemonkeyQueryLangLexer(input antlr.CharStream) *ImagemonkeyQueryLangLexer {

	l := new(ImagemonkeyQueryLangLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ImagemonkeyQueryLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ImagemonkeyQueryLangLexer tokens.
const (
	ImagemonkeyQueryLangLexerSEP                        = 1
	ImagemonkeyQueryLangLexerANNOTATION_COVERAGE_PREFIX = 2
	ImagemonkeyQueryLangLexerIMAGE_WIDTH_PREFIX         = 3
	ImagemonkeyQueryLangLexerIMAGE_HEIGHT_PREFIX        = 4
	ImagemonkeyQueryLangLexerPERCENT                    = 5
	ImagemonkeyQueryLangLexerPIXEL                      = 6
	ImagemonkeyQueryLangLexerOPERATOR                   = 7
	ImagemonkeyQueryLangLexerASSIGNMENT                 = 8
	ImagemonkeyQueryLangLexerORDER_BY                   = 9
	ImagemonkeyQueryLangLexerORDER_BY_VALIDATION_DESC   = 10
	ImagemonkeyQueryLangLexerORDER_BY_VALIDATION_ASC    = 11
	ImagemonkeyQueryLangLexerORDER_BY_VALIDATION        = 12
	ImagemonkeyQueryLangLexerLABEL                      = 13
	ImagemonkeyQueryLangLexerUUID                       = 14
	ImagemonkeyQueryLangLexerVAL                        = 15
	ImagemonkeyQueryLangLexerAND                        = 16
	ImagemonkeyQueryLangLexerOR                         = 17
	ImagemonkeyQueryLangLexerNOT                        = 18
	ImagemonkeyQueryLangLexerLPAR                       = 19
	ImagemonkeyQueryLangLexerRPAR                       = 20
	ImagemonkeyQueryLangLexerSKIPPED_TOKENS             = 21
)
