#!/usr/bin/env bash

cd testdata
rm b_hardlink.txt
ln b.txt b_hardlink.txt

cd dir2/dir3
rm d_hardlink.txt
ln ../../d.txt d_hardlink.txt
