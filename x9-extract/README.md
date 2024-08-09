# X9 Extract

The x9-extract tool extracts information from one or more [X9 files](https://www.frbservices.org/binaries/content/assets/crsocms/financial-services/check/setup/frb-x937-standards-reference.pdf) which can be used by the fin-ocr system to test or train OCR (Optical Character Recognition) of bank checks.

### Getting Started

1. [Install golang](https://go.dev/doc/install) if it is not already installed on your system.

2. Build the `x9-extract` binary as follows:

   ```
   go build
   ```

3. To see the usage message:

   ```
   $ x9-extract
   Usage: x9-extract <outputDir> <x9File1> [<x9File2> ...]
   ```

4. Preparing data

   It is always best to keep your testing and training data separate.

   In order to prepare your testing data, put the X9 files that you want to use for testing in your `$HOME/x9-files/test` directory and run the following command:

   ```
   FRB_COMPATIBILITY_MODE=true x9-extract $HOME/.fin-ocr/checks $HOME/x9-files/test/*
   ```

   This will populate your `$HOME/.fin-ocr/checks` directory with two files per check for testing: a TIFF file and a JSON file containing the necessary fields extracted from your X9 file.

   In order to prepare your training data, put the X9 files that you want to use for training in your `$HOME/x9-files/train` directory and run the following command

   ```
   FRB_COMPATIBILITY_MODE=true x9-extract $HOME/.fin-ocr/train/checks $HOME/x9-files/train/*
   ```

   This will populate your `$HOME/.fin-ocr/train/checks` directory with two files per check as described above.