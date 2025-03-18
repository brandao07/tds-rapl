#!/bin/bash

IGNORE_PROGRAMS=()
INCLUDE_PROGRAMS=()
IGNORE_LANGUAGES=()

# Initialize measurements file
echo "language,program,duration,emissions,energy_consumed" > codecarbon_measurements.csv

# Iterate over programs
for language in "Languages"/*; do
	language_name=$(basename "$language")
	# Check if the language is in the ignore list
	if printf '%s\n' "${IGNORE_LANGUAGES[@]}" | grep -qx "$language_name"; then
		echo "Ignoring language: $language_name"
		continue
	fi
	for program in "$language"/*; do
		if [ -d "$program" ]; then
			program_name=$(basename "$program")

			# Check if the program should be included or excluded
			if [ ${#INCLUDE_PROGRAMS[@]} -ne 0 ]; then
				if ! printf '%s\n' "${INCLUDE_PROGRAMS[@]}" | grep -qx "$program_name"; then
					echo "Skipping program not in include list: $program_name"
					continue
				fi
			else
				if printf '%s\n' "${IGNORE_PROGRAMS[@]}" | grep -qx "$program_name"; then
					echo "Ignoring program: $program_name"
					continue
				fi
			fi

			makefile_path="$program/Makefile"
			if [ -f "$makefile_path" ]; then
				cd $program
				echo "language,program,duration,emissions,energy_consumed" > codecarbon_measurements.csv
				make compile
				make codecarbon

				# Append results to measurements.csv
				file="codecarbon_measurements.csv"
				tail -n +2 "$file" >> ../../../codecarbon_measurements.csv
				make clean
				cd ../../..
			else
				echo "Makefile not found: $makefile_path"
			fi
		fi
	done
done