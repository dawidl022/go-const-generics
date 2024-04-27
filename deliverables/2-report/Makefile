main.pdf: project-plan.pdf *.tex examples/* appx/*
	pdflatex -shell-escape main
	biber main
	pdflatex -shell-escape main

project-plan.pdf: project-plan.mmd
	mmdc -i project-plan.mmd -o project-plan.pdf
