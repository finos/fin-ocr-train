# X9 Extract

X9 Extract is a tool for extracting check images and metadata from X9 files, which can be used to prepare data for training or testing Optical Character Recognition (OCR) systems for bank checks, particularly for MICR (Magnetic Ink Character Recognition) line recognition.

## Overview

- Extracts check images and metadata from X9 files
- Generates TIFF images and corresponding JSON files for each check

## Getting Started

### Prerequisites

- Go (version 1.21 or later)

### Installation

1. Build the `x9-extract` binary:

   ```
   go build
   ```

### Usage

```
x9-extract <outputDir> <x9File1> [<x9File2> ...]
```

- `<outputDir>`: Directory where extracted files will be saved
- `<x9File1>`, `<x9File2>`, etc.: One or more X9 files to process

### Example

Note: Setting the `FRB_COMPATIBILITY_MODE` environment variable to `true` enables Federal Reserve Bank (FRB) compatibility mode. This mode adjusts the processing of X9 files to meet specific FRB requirements. This mode allows certain fields to have default values or prevents errors that would occur in non-FRB contexts, ensuring compatibility with FRB standards. For example, this mode changes the DigitalSignatureMethod field from "0" to "00" when required. See (moov-io/imagecashletter repo)[https://github.com/search?q=repo%3Amoov-io%2Fimagecashletter%20IsFRBCompatibilityModeEnabled&type=code] to dig out the specifics if needed.

Prepare testing data:

```bash
FRB_COMPATIBILITY_MODE=true ./x9-extract $HOME/.fin-ocr/checks $HOME/x9-files/test/*
```

Prepare training data:

```bash
FRB_COMPATIBILITY_MODE=true ./x9-extract $HOME/.fin-ocr/train/checks $HOME/x9-files/train/*
```

## Output

For each check in the input X9 file(s), the tool generates:

1. A TIFF file containing the check image
2. A JSON file containing metadata extracted from the X9 file

Output files are named using the pattern `check-<num>.tiff` and `check-<num>.json`, where `<num>` is a monotonically increasing number.

### JSON Structure

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

## Vulnerability Report

To generate a report containing any vulnerabilities in any dependency please use govulncheck.

```bash
go install golang.org/x/vuln/cmd/govulncheck@latest
```

```bash
cd x9-extract
$HOME/go/bin/govulncheck ./...
```

## License Report

```bash
go install github.com/google/go-licenses@latest
```

```bash
$HOME/go/bin/go-licenses report github.com/finos/fin-ocr-train/... 
```
