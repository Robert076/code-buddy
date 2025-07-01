# 🌎 code-buddy

CLI tool for fetching the latest gitignore templates for your language when you start a new project. The templates are taken from the official [github gitignore repository](https://github.com/github/gitignore). As of right now, 155 gitignore templates are supported.

## 🍿 Live demo
https://github.com/user-attachments/assets/16e40405-fce4-4d0f-a1cd-a895b1041262

---

## 🚀 Setup
1. Clone the repository
```bash
git clone https://github.com/Robert076/code-buddy.git
```
2. Build the binary (you can also use go install instead)
```bash
cd code-buddy

cd cmd/code-buddy

go build -o cb
```
3. Take that binary file, and move it to `usr/local/bin` (this is for macOS, add it to your path on windows)
```bash
mv cb <your_path_to_bin>
```

4. Add executable permissions to binary
```bash
chmod +x cb
```

---

## 🧩 Use cases

1. Get the gitignore you want on your terminal window
```bash
cb gitignore <your_language>
``` 
2. Get the gitignore in your clipboard, ready to paste
```bash
cb gitignore <your_language> | pbcopy
```
3. Get the gitignore straight into a file (most likely `.gitignore`)
```bash
cb gitignore <your_language> >> (path_to .gitignore)
```
