use std::{
    env, fs, io::{self, Write}, path::PathBuf
};

fn main() {
    let roots_env = env::var("ProjectRoots").expect("ProjectRoots environment variable not set");
    let roots: Vec<&str> = roots_env.split(',').collect();

    let mut project = std::env::args().nth(1).or(Some("".to_string())).unwrap();
    if project.is_empty(){
        println!("输入要打开的项目的关键字：");
        io::stdout().flush().unwrap();
    
        io::stdin()
            .read_line(&mut project)
            .expect("Failed to read line"); 
    }

    let mut matched_projects = Vec::new();
    let matched: &str = project.trim().as_ref(); 
    for path in roots {
        // 遍历目录
        for entry in fs::read_dir(path).unwrap() {
            let entry = &entry.unwrap();
            let path = entry.path();
            if entry.file_name().into_string().unwrap().contains(matched) {
                matched_projects.push(path);
            }
        }
    }

    if matched_projects.is_empty() {
        println!("未匹配任何项目");
        return;
    }

    if matched_projects.len() == 1 {
        open_pro(&matched_projects[0]);
        return;
    }

    println!("匹配的项目为:");
    for (index, path) in matched_projects.iter().enumerate() {
        println!("{}. {}", index + 1, path.display());
    }

    println!("输入要打开的项目的索引号：");
    io::stdout().flush().unwrap();

    let mut input = String::new();
    io::stdin()
        .read_line(&mut input)
        .expect("Failed to read line");
    let choice = input
        .trim()
        .parse::<usize>()
        .expect("Please enter a valid number");

    if choice > 0 && choice <= matched_projects.len() {
        let selected_project = &matched_projects[choice - 1];
        open_pro(selected_project);
    } else {
        println!("Invalid choice. Please enter a valid number.");
    }
}

fn is_go(path: &PathBuf) -> bool {
    if path.join("go.mod").exists() {
        return true;
    }
    false
}

fn start_goland(path: &PathBuf) {
    let output = std::process::Command::new("goland")
        .arg(".")
        .current_dir(path)
        .output()
        .expect("无法启动goland");
    println!("{}", String::from_utf8_lossy(&output.stdout));
}

fn is_py(path: &PathBuf) -> bool {
    path.join("requirements.txt").exists()
}

fn is_rust(path: &PathBuf) -> bool {
    path.join("Cargo.toml").exists()
}

fn is_flutter(path: &PathBuf) -> bool {
    path.join("pubspec.yaml").exists()
}

fn start_vscode(path: &PathBuf) {
    let output = std::process::Command::new("code")
        .arg(".")
        .current_dir(path)
        .output()
        .expect("无法启动vscode");
    println!("{}", String::from_utf8_lossy(&output.stdout));
}

fn open_pro(path: &PathBuf) {
    println!("正在打开项目: {}", path.display());
    if is_go(path) {
        start_goland(path);
        return;
    }
    if is_py(path) || is_rust(path) || is_flutter(path) {
        start_vscode(path);
        return;
    }
    println!("未能识别出项目类别，请完善此程序，path:{}", path.display());
}
