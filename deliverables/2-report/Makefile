main.pdf: project-plan.pdf *.tex examples/*
	pdflatex main

project-plan.pdf: project-plan.mmd
	mmdc -i project-plan.mmd -o project-plan.pdf
