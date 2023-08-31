
from git import Repo

def git_diff():
    # 初始化 Repo 对象
    repo = Repo('.')

    # 获取最后一次提交
    last_commit = repo.head.commit

    # 获取这次提交和上一次提交之间的差异
    diffs = last_commit.diff('HEAD~1')

    # 遍历每个差异
    for diff in diffs.iter_change_type('M'):
        print(f'Changed: {diff.a_path}')

git_diff()