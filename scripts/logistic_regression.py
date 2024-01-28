import sys
import os
import pandas as pd
from sklearn.linear_model import LogisticRegression
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score

def perform_logistic_regression(csv_file, *custom_args):
    # Use custom_args as needed
    print("Custom Arguments:", custom_args)

    # Get the root directory of the project (assuming the Python script is in a subdirectory)
    root_dir = os.path.dirname(os.path.abspath(__file__))

    # Get the absolute path of the CSV file relative to the root directory
    csv_file_path = os.path.join(root_dir, csv_file)

    # Load the dataset (assuming a CSV format)
    df = pd.read_csv(csv_file_path)

    # Assuming the last column is the target variable and the rest are features
    X = df.iloc[:, :-1]
    y = df.iloc[:, -1]

    # Split the dataset into training and testing sets
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

    # Create a logistic regression model
    model = LogisticRegression()

    # Train the model
    model.fit(X_train, y_train)

    # Make predictions on the test set
    y_pred = model.predict(X_test)

    # Print accuracy score
    accuracy = accuracy_score(y_test, y_pred)
    print("Accuracy:", accuracy)

    # Invoke the database interaction callback
    if "--db-interaction" in custom_args:
        db_interaction_index = custom_args.index("--db-interaction")
        db_interaction_args = custom_args[db_interaction_index + 1:]
        invoke_db_interaction(*db_interaction_args)

    # Return the accuracy for further use
    return accuracy

def invoke_db_interaction(*db_args):
    # Use db_args to establish a database connection and perform operations
    print("Performing database interaction...")
    # Example: Connect to the database and execute queries

if __name__ == "__main__":
    # Example usage
    accuracy = perform_logistic_regression("../input_data.csv", *sys.argv[1:])
    print("Accuracy returned:", accuracy)
