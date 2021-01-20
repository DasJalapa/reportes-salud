// const excel = require('read-excel-file')

// excel('C:/Users/marol/Desktop/fonts/data.xlsx').then(rows => {
//     console.log(rows)
//     console.table(rows)
// }).catch(e => { console.log(e) })

const excel = require('xlsx')
const file = excel.readFile('./data.xlsx')
const uuid = require('uuid')
const job = require('./jobs')

let data = []

const sheet = file.SheetNames
for (let i = 0; i < sheet.length; i++) {
    const temp = excel.utils.sheet_to_json(file.Sheets[file.SheetNames[i]])

    temp.forEach(row => {
        data.push(row)
    })
}

data.forEach(element => {
    if(element['DEPENDENCIA DONDE LABORA']){
        console.log(`INSERT INTO workdependency VALUES ('${uuid.v4()}','${element['DEPENDENCIA DONDE LABORA']}');`)
    }
});

// let uuidjob;
// data.forEach(reg => {
//     // uuidjob = Object.keys(job)
//     var puesto = reg['Título del puesto']
//     var nombre = reg['NOMBRE COMPLETO DEL SERVIDOR']
//     var cui = reg['CÓDIGO DE EMPLEADO'] || null
//     uuidjob = job[puesto] || null


//     console.log(`INSERT INTO person (uuid, fullname, cui, job_uuid) VALUES('${uuid.v4()}', '${nombre}', '${cui}', '${uuidjob}');`)
// });


// let jobs = []

// data.forEach(element => {
//     if (!jobs.includes(element['Título del puesto'])) {
//         jobs.push(element['Título del puesto'])
//     }
// })


// let jobuuid = []
// let uuidjob;
// let registerjob;
// jobs.forEach(job => {
//     uuidjob = uuid.v4()
//     if (job) {
//         registerjob = {
//             uuid: uuidjob,
//             job: job
//         }

//         jobuuid.push(registerjob)
//         console.log(`INSERT INTO job VALUES ('${uuidjob}', '${job}', null);`)
//     }
// })

// console.log('------------------------')
// console.log(jobuuid)