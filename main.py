import json
import os
import sys
import openai

from gitxxx import git_diff


# 获取环境变量 OPENAI_API_KEY
def get_openai_key():
    key = os.getenv('OPENAI_API_KEY')
    if not key:
        print("not found OPENAI_API_KEY")
        sys.exit()
    return key


openai.api_key = get_openai_key()

sys_msg = '''
As a software development assistant, your role is to generate commit message for provided code diff. Please adhere to the following guidelines:

1. Commit message should consist of a concise title and ensure that the title is within 50 characters in length.
2. The code diff will be in JSON format, comprising an array with three fields: file, added, and removed.
3. Prioritize actual code changes and disregard any irrelevant modifications.
'''



if __name__ == '__main__':
    diff_output = git_diff()
    print(json.dumps(diff_output))
    if len(diff_output) == 0:
        print("noting need to commit")
        sys.exit()

    completion = openai.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "system", "content": sys_msg},
            {"role": "user", "content": json.dumps(diff_output)}
        ]
    )

    print(completion.choices[0].message)
    print(completion)

'''mine
It seems like you have provided a Python script that uses the `unidiff` library to parse and analyze differences in a Git repository.
The script consists of a function `git_diff()` that returns a list of dictionaries containing information about the modified files.
\n\nTo summarize the functionality of the script:\n\n
1. Import the necessary `unidiff` module: `from unidiff import PatchSet`.\n\n
2. Create a `Repo` object representing the current repository: `repo = Repo()`.\n\n
3. Get the `git diff` output as a string: `diffs = repo.git.diff()`.\n\n
4. Parse the diff string into a `PatchSet` object: `patch_set = PatchSet(diffs)`.\n\n
5. Iterate over each patch in the `PatchSet` object:\n\n 
  - Create a dictionary `file` to store information about the file with keys 'file', 'added', and 'removed'.\n\n
  - Iterate over each hunk in the patch:\n\n     
  - Iterate over each line in the hunk:\n\n       
  - Skip empty lines and lines starting with '#' characters.\n\n   
  - If the line is added, append its value to the 'added' field in the `file` dictionary.\n\n 
  - If the line is removed, append its value to the 'removed' field in the `file` dictionary.\n\n  
  - Append the `file` dictionary to the `files` list.\n\n
6. Return the `files` list.\n\n
7. Outside the `git_diff()` function, iterate over each file dictionary in the `files` list and print the file path, added lines, and removed lines.\n\n
However, it seems like there are some parts of the code that are commented out or have non-ASCII charactershhgfhjdh65为其2 x If you need further assistance, please provide more information or clarify your intentions.
'''

'''devchat
Refactor git_diff function to parse git diff output

- Replace Repo('.') with Repo() to represent the current repository.
- Replace the method of getting diffs. Now it gets git diff output directly instead of getting the difference between the last commit and the previous one.
- Use PatchSet from unidiff library to parse the git diff output.
- Add logic to skip empty lines and lines starting with '#'.
- Store the parsed information (file name, added content, removed content) in a list of dictionaries.
- Print out the parsed information instead of just printing the changed file paths.
'''