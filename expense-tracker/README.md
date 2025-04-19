# Expense Tracker (CLI - Go)

A **Command-Line Expense Tracker** written in **Go**, designed to help you quickly log and analyze expenses from your terminal. Expenses are stored in a **CSV** file, making it lightweight, portable, and easy to back up.

Inspired by the [roadmap.sh Expense Tracker project](https://roadmap.sh/projects/expense-tracker).

---

## Features

- Add a new expense with a description, amount, and date  
- Update an existing expense by ID  
- Delete one or more expenses by ID  
- List all recorded expenses  
- View a **summary of expenses** for:
  - A specific **day**
  - A specific **month**
  - A specific **year**
  - **All time**  
- Export listed expenses to a CSV file using the `--export` flag  
- Generate shell autocompletion scripts  
- Customizable file paths via flags or environment defaults

---

## Installation

### Clone and Build

```bash
git clone https://github.com/idukrystal/Expense-Tracker.git
cd Expense-Tracker/expense-tracker
go build -o expense

Make the CLI Globally Accessible

To make the expense command available globally, move the compiled binary to a directory that is included in your system's PATH.

Option 1: Move to /usr/local/bin/ (system-wide, requires sudo)

sudo mv expense /usr/local/bin/

Option 2: Move to ~/bin/ (user-specific, no sudo needed)

1. Create the ~/bin/ directory if it doesn't exist:

mkdir -p ~/bin


2. Move the binary to ~/bin/:

mv expense ~/bin/


3. Add ~/bin to your PATH if it's not already there:

Open your shell configuration file (~/.bashrc, ~/.zshrc, or ~/.bash_profile) and add the following line:

export PATH="$HOME/bin:$PATH"


4. Reload the shell configuration:

source ~/.bashrc
# or
source ~/.zshrc



Verify the Global Availability

Once the binary is placed in a directory within your PATH, you can run the expense command from any directory in the terminal:

expense --help

This should show you the help message, confirming that the app is available globally.


---

Usage

expense {<operation> [--<key> <value> | --help] ... | --help} [flags]
expense [command]

Available Commands


---

Example Commands

Add an Expense

expense add --description "Lunch" --amount 2500 --day 19 --month 4 --year 2025

Update an Expense

expense update --id 2 --description "Dinner" --amount 3000

Delete an Expense

expense delete --id 2

List All Expenses

expense list

Export Expenses to CSV

expense list --export /path/to/exported_expenses.csv

You can specify any valid file path to export the list of expenses. The export feature works directly with the list command.

View Summary

# All-time summary
expense summary

# Daily summary
expense summary --day 19 --month 4 --year 2025

# Monthly summary
expense summary --month 4 --year 2025

# Yearly summary
expense summary --year 2025


---

Flags


---

CSV Format

Expenses are stored in a CSV file (expenses.csv) with this format:

id,date,description,amount
1,2025-04-19,Lunch,2500
2,2025-04-20,Transport,1000


---

Future Improvements

Tag and categorize expenses

Range-based filtering (e.g., between dates)

Visual summaries or CLI bar charts

Encrypted storage

Cloud sync option



---

License

MIT License — see LICENSE for details.


---

Now, after following these instructions, the expense command will be available globally on your system. You can run it from any directory in the terminal, and it will work as expected!

Let me know if you need further adjustments or have any questions.

This is now in markdown format, ready to be used as your `README.md` file. Let me know if you need any further changes!

