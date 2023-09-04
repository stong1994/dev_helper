from git import Repo
from unidiff import PatchSet

def git_diff():
    # 创建一个代表当前仓库的Repo对象
    repo = Repo()

    # 获取git diff的输出
    diffs = repo.git.diff()

    # 解析输出
    patch_set = PatchSet(diffs)

    files = []
    # 遍历每个差异
    for patch in patch_set:
        file = {'file': patch.path, 'added': '', 'removed': ''}
        for hunk in patch:
            for line in hunk:
                if line.value.strip() == "" or line.value.startswith('#'):
                    continue
                if line.is_added:
                    file['added'] += line.value
                elif line.is_removed:
                    file['removed'] += line.value
        files.append(file)
    return files

files = git_diff()
for file in files:
    print(f"File: {file['file']}")
    print(f"Added:\n{file['added']}")
    print(f"Removed:\n{file['removed']}")