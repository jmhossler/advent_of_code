def read_numbers(file_name):
    return transform_file(file_name, int)

def read_strings(file_name):
    return transform_file(file_name, str)

def transform_file(file_name, transform):
    with open(file_name) as f:
        return [transform(row.strip()) for row in f.readlines()]
