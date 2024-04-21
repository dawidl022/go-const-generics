main.pdf: project-plan.pdf *.tex examples/*
	pdflatex -shell-escape main
	biber main
	pdflatex -shell-escape main

project-plan.pdf: project-plan.mmd
	mmdc -i project-plan.mmd -o project-plan.pdf
