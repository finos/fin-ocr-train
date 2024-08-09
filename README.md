# Training tesseract

This repository contains information about what has been done to train tesseract for the MICR line of bank checks; however, the same general approach can apply to other fonts on other media.

Two different methods for training tesseract have been used as follows:

1. Using real checks - see [real](./real/README.md).

   The best results have been achieved with this method so far.

2. Using synthetically generated MICR lines of checks - see [synthetic](./synthetic/README.md).

   This appears to be what most people in the community are doing; however, thus far, the resulting accuracy has been significantly lower using this method.
