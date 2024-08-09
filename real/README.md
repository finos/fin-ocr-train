# MICR training with real checks

### Prerequisites

1. Install python (e.g. `brew install python`)

2. Install pillow (e.g. `brew install pillow`)

### How to create training data

#### Creating training data automatically

If you have one or more X9 files, see [x9-extract](../x9-extract/README.md) to automate the creation of your training data.  Keep in mind that some of the data in the X9 files may not be correct; therefore, you may need to manually correct some of the values in the resulting JSON files.  The quality of your training depends on the accuracy of the values in the JSON files.

#### Creating training data manually

This section describes how to manually create your training data files.  If you do not have any X9 files or just want to understand the format of your training data, this will be useful.

All training data must be in a single directory.  This directory must contain 
two files for each check to use in training: `<prefix>.tiff` and `<prefix>.json`.  The tiff file is the image of the check in TIFF format, and
the JSON file contains the correct data from the MICR line of the check in the following format:

```
{
   "auxiliaryOnUs": "1234567",
   "payorBankRoutingNumber": "12345678",
   "payorBankCheckDigit": "1",
   "onUs": "1234567890/"
}
```

### How to start a training session

In order to start a training session, run the following command:

```
mgr train <starting-num> <count>
```

For example, if you have 20000 pairs of training input files, run:

```
mgr train 1 20000
```

This is a very CPU intensive process and can take many hours to run; therefore, you may want to run it in the background using a utility such as the unix `nohup` command.

When processing is complete, the results are stored in a sub-directory of `$HOME/.fin-ocr/train/results`.  The name of the sub-directory is based on the current date.  Two files are stored in that sub-directory: 
* `micr_e13b.traineddata` - the resulting traineddata file which can now be used by the [CLI](https://github.com/discoverfinancial/fin-ocr-cli) or [REST service](https://github.com/discoverfinancial/fin-ocr-web);
* `train.log` - the logs from the training session.