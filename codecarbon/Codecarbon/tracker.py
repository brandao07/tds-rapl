import os
import csv
import sys
from codecarbon import OfflineEmissionsTracker


program = sys.argv[1]
term = sys.argv[2]
language = sys.argv[3]
program_name = sys.argv[4]
num_executions = int(sys.argv[5])

csv_filename = "codecarbon_measurements.csv"

for execution in range(1, num_executions + 1):
    tracker = OfflineEmissionsTracker(country_iso_code="PRT")

    print(f"Running execution {execution} of {num_executions}...")

    tracker.start()

    if language == "C":
        os.system(f"./{program} {term}")
    elif language == "JavaScript":
        os.system(f"node {program} {term}")

    tracker.stop()

    data = tracker.final_emissions_data
    row = [language, program_name, data.duration, data.emissions, data.energy_consumed]

    with open(csv_filename, mode="a", newline="") as file:
        writer = csv.writer(file)
        writer.writerow(row)