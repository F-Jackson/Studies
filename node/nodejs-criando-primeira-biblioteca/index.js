import chalk from "chalk";
import fs from "fs";

function trataErro(err) {
    throw new Error(chalk.red(err.code, 'nao ha arquivo no diretorio'));
}

function pegaArquivo(caminho) {
    const encoding = 'utf-8';
    fs.promises
        .readFile(caminho, encoding)
        .then((text) => console.log(text))
        .catch((err) => trataErro(err))
}

pegaArquivo('./arquivos/texto.md')