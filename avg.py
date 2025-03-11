import csv

def calculate_avg_without_min_max(values):
    if len(values) <= 2:
        raise ValueError("List must have more than two elements to remove min and max values.")
    
    values_sorted = sorted(values)
    values_filtered = values_sorted[1:-1]  # Remove min and max
    
    avg_value = sum(values_filtered) / len(values_filtered)
    return avg_value

def read_csv_column(file_path, column_name):
    values = []
    with open(file_path, 'r') as file:
        reader = csv.DictReader(file)
        for row in reader:
            values.append(float(row[column_name]))
    return values

csv_file_path = '/path'  
package_values = read_csv_column(csv_file_path, ' Package')
time_values = read_csv_column(csv_file_path, ' Time (ms)')

avg_package = calculate_avg_without_min_max(package_values)
avg_time = calculate_avg_without_min_max(time_values)

print(f"Average Energy Consumption: {avg_package:.2f} J")
print(f"Average Time: {avg_time:.2f} ms")
