use std::{
    fs::{self, DirEntry},
    path::PathBuf,
};

fn main() {
    // 获取命令行参数
    // let path = std::env::args().nth(1).expect("路径参数不能为空");
    let project = "ucenter";

    let roots = [
        "/Users/stong/Project/Personal",
        "/Users/stong/Project/Company",
        "/Users/stong/Project/Github",
    ];

    for path in roots {
        // 遍历目录
        for entry in fs::read_dir(path).unwrap() {
            let entry = &entry.unwrap();
            let path = entry.path();

            if !entry.file_name().into_string().unwrap().contains(project){
                continue;
            }

            println!("运行 cargo run 在 {}", path.display());
            if is_go(entry) {
                start_goland(path);
                return;
            }
            if is_py(entry) || is_rust(entry) || is_flutter(entry) {
                start_vscode(path);
                return;
            }
            println!("未能识别出项目类别，请完善此程序，path:{}", path.display());
        }
    }
}

fn is_go(dir: &DirEntry) -> bool {
    let path = dir.path();
    if path.join("go.mod").exists() {
        return true;
    }
    false
}

fn start_goland(path: PathBuf) {
    let output = std::process::Command::new("goland")
        .arg(".")
        .current_dir(path)
        .output()
        .expect("无法启动goland");
    println!("{}", String::from_utf8_lossy(&output.stdout));
}

fn is_py(dir: &DirEntry) -> bool {
    dir.path().join("requirements.txt").exists()
}

fn is_rust(dir: &DirEntry) -> bool {
    dir.path().join("Cargo.toml").exists()
}

fn is_flutter(dir: &DirEntry) -> bool {
    dir.path().join("pubspec.yaml").exists()
}

fn start_vscode(path: PathBuf) {
    let output = std::process::Command::new("code")
        .arg(".")
        .current_dir(path)
        .output()
        .expect("无法启动vscode");
    println!("{}", String::from_utf8_lossy(&output.stdout));
}
