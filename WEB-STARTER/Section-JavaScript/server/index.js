const express = require('express')
const bodyparser = require('body-parser')
const app = express()

app.use(bodyparser.json())

const port = 8000

// users
let users = []
let counter = 1

// part = GET /
app.get('/users', (req, res) => {
    const filterUser = users.map(user => {
        return {
            id: user.id,
            firstname: user.firstname,
            lastname: user.lastname,
            fullname: user.firstname + ' ' + user.lastname
        }
    })
    res.json(filterUser)
})

// part = POST /user
app.post('/users', (req, res) => {
    let user = req.body
    user.id = counter
    counter += 1

    users.push(user)
    res.json({
        message: 'add ok',
        user: user
    })
})

// // path =  GET /users/:id สำหรับการดึง users รายคนออกมา
// app.get('/users/:id', (req, res) => {
//     let id = req.params.id

//     //find index od user by id
//     let selectedIndex = users.findIndex(user => user.id == id)

//     // find index 
//     res.json(users[selectedIndex])
// })

// // path = PUT /users/:id
// app.put('/users/:id', (req, res) => {
//     let id = req.params.id
//     let updateUser = req.body

//     // find user by id 
//     let selectedIndex = users.findIndex(user => user.id == id)
    
//     // then update user with new data (null || test)
//     users[selectedIndex].firstname = updateUser.firstname || users[selectedIndex].firstname
//     users[selectedIndex].lastname = updateUser.lastname || users[selectedIndex].lastname
//     users[selectedIndex].age = updateUser.age || users[selectedIndex].age
//     users[selectedIndex].gender = updateUser.gender || users[selectedIndex].gender

//     res.json({
//         message: 'update user complete',
//         data: {
//             user: updateUser,
//             indexUpdate: selectedIndex
//         }
//     })
// })

// path = PATCH /users/:id
app.patch('/users/:id', (req, res) => {
    let id = req.params.id
    let updateUser = req.body

    // find user by id 
    let selectedIndex = users.findIndex(user => user.id == id)
    
    // then update user with new data (null || test)
    if (updateUser.firstname) {
        users[selectedIndex].firstname = updateUser.firstname
    }

    if (updateUser.lastname) {
        users[selectedIndex].lastname = updateUser.lastname
    }

    res.send({
        message: 'update user complete',
        data: {
            user: updateUser,
            indexUpdate: selectedIndex
        }
    })
})


// path = DELETE /users/:id
app.delete('/users/:id', (req, res) => {
    let id = req.params.id

    // find index of user by id
    let selectedIndex = users.findIndex(user => user.id == id)

    // // delete
    // delete users[seletedIndex]

    // delete user
    users.splice(selectedIndex, 1)
    
    res.json({
        message: 'delete complete!',
        indexDeleted: selectedIndex
    })
})

app.listen(port, (req, res) => {
    console.log('Server is ranning on port' + port)
})