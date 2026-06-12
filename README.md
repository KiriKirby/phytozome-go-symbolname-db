# phytozome GO symbolname DB

This repository builds and publishes the prebuilt `symbolname.pgd` database for phytozome GO.

The scheduled workflow reads the full NCBI Gene `GENE_INFO/` split source directory, builds a compact `.pgd`, compresses it, publishes archive parts as GitHub Release assets, and updates the manifest branch only after all assets upload successfully.

Main application code lives in `KiriKirby/phytozome-go`; this repository is intentionally only for database build and distribution artifacts.
