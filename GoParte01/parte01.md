Parte 01
pré requisitos:
    - computador(8gm ram, bom processador, ssd 500gb)

Aqui está como fazer a primeira parte:
-Instalar o vscode. (site)
- Aprender a versionar em github. (site)
- foco em back end - iniciar aprender o básico da estrutura de arquivos do go
- não use ia ainda. 

Preparo de ambiente:
- instalar Go no vscode
- playground do go((https://go.dev/play/))

Primeiros passos:

comandos para verificar:
- windows
 - abra o cmd
 - vamos usar o Chocolatey como gerencador de pacotes pra gente
 -abra o cmd e cole o comando: 
 @"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"
 
 - para testar digite: choco
 instalar o go: choco install go   (necessário ao go)
 instalar o mingw: choco install mingw -y  (necessário se usar c)(opcional)
 - digite: go version
 - vai aparecer algo como: "go version go1.25.1 windows/amd64"
 -instalando o git: - choco install git -y
-instalando o github: - choco install gh -y
 - Extra - criar ssh para conectar no github
         - ssh-keygen -t ed25519 -C "seu_email_do_github"
         - eval $(ssh-agent -s)
         --saber qual user digite: whoami
         - ssh-add $env:USERPROFILE\.ssh\id_ed25519
         - Get-Content $env:USERPROFILE\.ssh\id_ed25519.pub
         - vai aparecer algo como: ssh-ed25519 AAAAC3...seutoken... seunome@pc
         4. Adicionar no GitHub

Vá em https://github.com/settings/keys

Clique em New SSH Key

Cole a chave no campo

Nomeie como "Windows-PC"

Se deu certo, aparece algo como:
Hi JeanMRocha! You've successfully authenticated, but GitHub does not provide shell access.

git remote set-url origin git@github.com:JeanMRocha/exercicios-go.git
daqui pra frente vai funcionar sem pedir senha

Criar uma pasta para iniciar o primeiro projeto. 



Para enviar para o github: 
Na aba Terminal do VS Code (que é igual ao CMD/PowerShell), os comandos são:

git add .
git commit -m "sua mensagem de commit"
git push


git add . → adiciona todas as mudanças.

git commit -m "mensagem" → cria o commit.

git push → envia para o GitHub.

