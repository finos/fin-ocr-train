# MICR Line Training for Tesseract OCR

This repository contains methods and tools for training Tesseract OCR to recognize MICR (Magnetic Ink Character Recognition) lines on bank checks. While the focus is on MICR, the general approach can be applied to other specialized fonts and media.

## Overview

The project explores two main approaches for training Tesseract:

1. Using real check images
2. Using synthetically generated MICR lines

Each method has its own advantages and challenges, which are detailed in their respective directories.

## Approaches

### 1. Real Check Training

Located in the `real/` directory, this method uses actual check images for training. Currently, this approach has yielded the best results in terms of accuracy.

[Learn more about real check training](./real/README.md)

### 2. Synthetic Data Training

Found in the `synthetic/` directory, this method generates artificial MICR lines for training. It has shown lower accuracy compared to using real checks in our tests.

[Learn more about synthetic data training](./synthetic/README.md)

## Additional Tools

### X9 Extract

The `x9-extract/` directory contains a tool for extracting check details from X9 files, which can be used to prepare data for training the OCR system.

[Learn more about X9 Extract](./x9-extract/README.md)

