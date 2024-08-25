from airflow import DAG
from airflow.operators.bash import BashOperator
from airflow.operators.python import PythonOperator
from datetime import datetime

def python_task():
    print("This is task 3")

with DAG(
    'DAG_STEPS',
    default_args={'owner': 'sdesdo'},
    description='DAG description',
    schedule_interval='0 0 * * *',
    start_date=datetime(2025, 1, 1)
) as dag:

    t1 = BashOperator(
        task_id='task_1',
        bash_command='echo "This is task 1"'
    )

    t2 = BashOperator(
        task_id='task_2',
        bash_command='echo "This is task 2"'
    )

    t3 = PythonOperator(
        task_id='task_3',
        python_callable=python_task
    )

    t1 >> t2 >> t3