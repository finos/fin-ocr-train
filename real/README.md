# MICR Training with Real Checks

This guide explains how to train Tesseract OCR for recognizing MICR (Magnetic Ink Character Recognition) lines on bank checks using real check images and the OCR CLI tool.

## Prerequisites

* **Node.js** and npm (for running the OCR CLI tool)
* **gmake** or **make** v4.2 or newer
* **Python 3.x** along with the `pillow` library for image processing.
* **wget**
* **Tesseract 5.x** (see [Tesseract installation instructions](https://tesseract-ocr.github.io/tessdoc/Installation.html))
* **OCR CLI tool** (follow installation instructions from the [fin-ocr-cli repository](https://github.com/finos/fin-ocr-cli))

## Preparing Training Data

You can prepare training data either automatically using X9 files or manually. 

### Option 1: Automatic Data Preparation (Using X9 Files)

If you have X9 files, you can use the [`x9-extract`](../x9-extract/README.md) tool to automate the creation of your training data.

1. Use [`x9-extract`](../x9-extract/README.md) to extract check images and metadata:
   ```
   FRB_COMPATIBILITY_MODE=true ./x9-extract $HOME/.fin-ocr/checks $HOME/x9-files/train/*
   ```

2. Use the OCR CLI to process and validate the extracted data:
   ```
   ocr check scan <start-check-num> <end-check-num>
   ```

   For example, to process checks 1 through 1000:
   ```
   ocr check scan 1 1000
   ```

   This command will:
   - Process each check image in the `$HOME/.fin-ocr/checks` directory
   - Compare the OCR results with the metadata from X9 files
   - Generate ground truth files for training

**Important Note:** Some data in the X9 files may not be accurate. The OCR CLI helps identify potential discrepancies by comparing OCR results with X9 metadata. You may need to manually review and correct values in the resulting JSON files to ensure training quality.

#### Ground Truth Generation

The OCR CLI is responsible in this flow for generating ground truth data, which is essential for Tesseract training. Here's what happens during this process:

1. For each check, the OCR CLI generates two important files 
   - `preprocessedImageFile` (TIFF): Contains the isolated MICR line image.
   - `groundTruthFile` (gt.txt): Contains the correct text corresponding to the MICR line.

Note: In some cases, ground truth files may not be created for a particular check. 

This can happen if:
 * The OCR results do not match the data in the JSON file

### Option 2: Manual Data Preparation

If you don't have X9 files or prefer manual setup, you can create the training data files yourself.

1. Create a directory for your training data: `$HOME/.fin-ocr/checks` or set the environment variable CHECKS_DIR to your preferred location.

2. For each check, create two files in this directory:
   - `check-<num>.tiff`: The check image in TIFF format
   - `check-<num>.json`: A JSON file containing the MICR line data

   JSON file schema:
```json
{
  "id": "check-<num>",
  "fileName": "original_x9_filename",
  "fileSeqNo": 1,
  "routingNumber": "123456789",
  "accountNumber": "1234567",
  "checkNumber": "1001",
  "auxiliaryOnUs": "1001",
  "payorBankRoutingNumber": "12345678",
  "payorBankCheckDigit": "9",
  "onUs": "1234567/1001"
}
```

3. You can use the OCR CLI to process your manually prepared data:
   ```
   ocr check scan <start-check-num> <end-check-num>
   ```

## Starting a Training Session

1. Ensure you're in the `real` directory of this repository.

2. Run the training command:
   ```
   ./mgr train <starting-num> <count>
   ```
   
   For example, to train on 20,000 checks:
   ```
   ./mgr train 1 20000
   ```

   **Note:** This process is CPU-intensive and may take several hours. Consider using `nohup` or a similar tool to run it in the background.

e.g.
```bash
nohup ./mgr train 1 20000 > training_output.log 2>&1 &
```
You can monitor the progress by checking the log file: `tail -f training_output.log`

Once the training is complete, you can find the training results in $HOME/.fin-ocr/train/results/<date>

## What Happens During Training

1. The `mgr` script sets up the training environment:
   - Clones the [tesstrain](https://github.com/tesseract-ocr/tesstrain) repository if not present
   - Downloads necessary language data files
   - Uses the OCR CLI to generate and validate ground truth data from your check images and JSON files

2. The script then initiates the Tesseract training process using the tesstrain make file.

## Handling Inaccurate X9 Data

If you encounter inaccuracies in the X9 data:

1. Review the OCR CLI output for any mismatches between OCR results and X9 metadata.
2. Manually inspect the check images and JSON files for discrepancies.
3. Update the JSON files with correct information if needed.
4. Consider creating a `corrections` file to track and apply corrections automatically. (TODO: expand on this)

## Training Results

After training completes:

1. Results are stored in `$HOME/.fin-ocr/train/results/<date>/`:
   - `micr_e13b.traineddata`: The trained Tesseract data file
   - `$HOME/.fin-ocr/train/train.log`: Detailed logs 

2. You can use the resulting `micr_e13b.traineddata` file with the [fin-ocr CLI](https://github.com/finos/fin-ocr-cli) or [fin-ocr REST service](https://github.com/finos/fin-ocr-rest) for MICR line recognition.

## Troubleshooting

- If training fails, check the `train.log` file for error messages (`$HOME/.fin-ocr/train/train.log` or `$TRAIN_DIR/train.log`)
- If using X9 files, verify that the extracted data is correct and complete. Use the OCR CLI to validate and identify potential issues.

## Additional Notes

- You may need to experiment with different training parameters (e.g., number of iterations, learning rates) to achieve optimal results for your specific dataset.
