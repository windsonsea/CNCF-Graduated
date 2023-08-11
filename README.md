# CNCF_Graduated
 - This repository includes all the CNCF graduated projects' documentation sites.
 - It is renewed everyday automatically.
 - The documentation sites do not exist as submodules, but as plain files.
 - The sites are included in the modules directory.
 - If you only want to clone the .md files, you can run:
  >git clone -n --depth=1 --filter=tree:0 https://github.com/365614269/CNCF_Graduated.git
  >cd CNCF_Graduated
  >git sparse-checkout set --no-cone '*.md'
  >git checkout
 - Having fun working on CNCF graduated documentation sites!