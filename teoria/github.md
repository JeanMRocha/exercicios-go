# 1. Cancelar qualquer rebase travado
git rebase --abort

# 2. Remover o arquivo de lock (se existir)
del ".git\index.lock"

# 3. Verificar o status do repositório
git status
