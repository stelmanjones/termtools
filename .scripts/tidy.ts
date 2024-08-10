import chalk from 'chalk'
async function tidyModules() {
    const modules = await Bun.$`go list -m`.lines()

    const names: string[] = []
    for await (const module of modules) {
        if (module !== "" && module !== "github.com/stelmanjones/termtools") {
            const name = module.split("/").reverse()[0]

            names.push(name)
        }

    }





    for (const name of names) {
        Bun.$.cwd(`./${name}`)
        Bun.$`go mod tidy`
        Bun.$.cwd("..")
        console.log(`Tidied up module ${chalk.green(name)}`)
    }

}


await tidyModules()

