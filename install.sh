#!/bin/bash

echo "Compilando o programa..."
go build

echo "Movendo o executável para /usr/bin..."
sudo mv goWallpaper /usr/bin

if [ $? -eq 0 ]; then
    echo "Instalação concluída com sucesso!"
    echo "Use o comando **goWallpaper** para executar o programa."
else
    echo "Ocorreu um erro durante a instalação."
fi
