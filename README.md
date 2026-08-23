# Task Tracker 
Sample solution for the task-tracker challenge from roadmap.sh.
## How to run
```cmd
git clone git@github.com:Kotoninja/task-tracker.git
cd task-tracker
go build -o task-cli
```
## The list of commands and their usage is given below:
```
# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```
## Package list
1. Cobra - library for creating powerful modern CLI applications
2. Tablewriter - library for generating rich text-based tables