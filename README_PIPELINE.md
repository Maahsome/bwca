# Pipeline Documentation

## Build-ChangeLog

This pipeline runs on a push/merge to main, it finds the last RELEASE, uses that to find MRs since the last TAG (release) and parses all of the MR `descriptions` for changelog info, and builds a v0.0.0-NEXT.md file, which it commits back to `changelog/v0.0.0-NEXT.md`.

