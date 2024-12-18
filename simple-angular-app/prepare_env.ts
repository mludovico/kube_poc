const setEnv = (env: String) => {
  const colors = require('colors');
  console.log(colors.magenta(`The current environment is: ${env}\n`));
  const fs = require('fs');
  const writeFile = fs.writeFile;
// Configure Angular `environment.ts` file path
  const targetFileName = `environment${env == 'prod' ? '' : '.dev'}.ts`;
  const targetPath = `./src/app/environments/${targetFileName}`;
// Load node modules
  const envFileName = `.env${env == 'prod' ? '' : '.dev'}`;
  console.log(colors.magenta(`The file ${envFileName} will be loaded to generate the ${targetFileName} file\n`));
  require('dotenv').config({
    path: envFileName
  });
// `environment.ts` file structure
  const envConfigFile = `export const environment = {
  production: ${env == 'prod' ? 'true' : 'false'},
  apiHost: '${process.env["API_HOST"]}',
  apiPort: '${process.env["API_PORT"]}',
};
`;
  console.log(colors.magenta('The file `environment.ts` will be written with the following content: \n'));
  writeFile(targetPath, envConfigFile, (err: any) => {
    if (err) {
      console.error(err);
      throw err;
    } else {
      console.log(colors.magenta(`Angular environment.ts file generated correctly at ${targetPath} \n`));
    }
  });
};

setEnv(process.env["npm_config_env"] || 'dev');
