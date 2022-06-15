const {exec,execSync} = require("child_process");
const path = require("path");
const fs = require("fs");

const program = "protoc";

const rootPath = typeof process.argv[1]  === "undefined" ? process.cwd() : process.argv[1];
const srcPath = typeof process.argv[2]  === "undefined" ? `${process.cwd()}/protobuf` : process.argv[2];
const outputPath = typeof process.argv[3]  === "undefined" ? `${process.cwd()}/pkgs` : process.argv[3];

const supportLang = [
    "rust",
    "go"
];

const makeOutPutFlag = (lang,value) => ` --${lang}_out=${value} `;
const IncludeDirFlag = (value) => ` -I ${value} `;
const joinPath = (root,sub) => `${root}/${sub}`;


const subNodePathsAction = (root) => fs.readdirSync(root).map(value => joinPath(root,value));
const popDirectoryPathAction = (list) => list.filter(value => fs.lstatSync(value).isDirectory());
const IncludeFlagListAction = (list) => list.map(value => IncludeDirFlag(value));

const sourceDirPathListAction = (root) => {
    let buf = [root];
    let index = 0;
    (()=> {
        while(index != buf.length) {
            let list = subNodePathsAction(buf[index]);
            buf.push(...popDirectoryPathAction(list));
            index++;
        }
    })();
    return buf;
}

const sourceFilePathListAction = (dir) => subNodePathsAction(dir).filter(value => fs.lstatSync(value).isFile());
const onlySourceFilePathListAction = (dir) => subNodePathsAction(dir)
                                                                                .filter(value => fs.lstatSync(value).isFile())
                                                                                .map(value => path.basename(value));
const makeIncludeFlagStrAction = (list) => list.map(value => IncludeDirFlag(value)).reduce((pre,value) => pre + value);
const printCompileResult = (err,stdout,stderr) => {
    if(err)throw err;
    if(stdout != "")console.log(stdout);
    if(stderr != "")console.log(stderr);
}

const compileProtobufAction = (iFlag,lang,outputDir,sourceFileList) => {
    exec(`protoc ${iFlag} ${makeOutPutFlag(lang,outputDir)} ${sourceFileList.reduce((pre,value) => `${pre} ${value}`)}`,printCompileResult);
}

const srcDirs = sourceDirPathListAction(srcPath);

const supportLangAction = (lang) => {
    srcDirs.forEach( value => {
        const filePaths = onlySourceFilePathListAction(value);
        if(filePaths.length == 0)return;
        if(!fs.existsSync(`${outputPath}/${lang}/protobuf`))fs.mkdirSync(`${outputPath}/${lang}/protobuf`);
        compileProtobufAction(makeIncludeFlagStrAction(srcDirs),lang,`${outputPath}/${lang}/protobuf`,filePaths);
    });
}
supportLang.forEach(supportLangAction);


