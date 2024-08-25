import chalk from 'chalk'
import { $ } from "bun";

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
        $.cwd(`./${name}`)
        const out = await $`go mod tidy`.nothrow().text()
        $.cwd("..")
        console.log(out)
        console.log(`Tidied up module ${chalk.green(name)}`)
    }

}

$.cwd("/Users/stelmanjones/Repositories/Go/termtools")

await tidyModules()

