const express = require('express')
const bodyparser = require('body-parser')
const mysql = require('mysql2/promise')
const cors = require('cors')
const app = express()

app.use(bodyparser.json())
app.use(cors())

const port = 8000

let conn = null

const initMySQL = async () => {
    conn = await mysql.createConnection({
        host: 'localhost',
        user: 'root',
        password: '',
        database: 'tutorials',
        port: 3306
    })
}

const validateData = (userData) => {
    let errors = []

    if (!userData.firstname) {
        errors.push('please input first name')
    }

    if (!userData.lastname) {
        errors.push('please input last name')
    }

    if (!userData.age) {
        errors.push('please input age')
    }

    if (!userData.gender) {
        errors.push('please input gender')
    }

    if (!userData.interests) {
        errors.push('please input interest')
    }

    if (!userData.description) {
        errors.push('please input description')
    }

    return errors
}

// part = GET /
app.get('/users', async (req, res) => {
    try {
        const results = await conn.query('SELECT * FROM users')
        res.json(results[0])
    } catch (error) {
        console.error('Error fetching users:', error.message)
        res.status(500).json({ error: 'Error fetching users' })
    }
})

// part = POST /user
app.post('/users', async (req, res) => {
    try {
        let user = req.body

        const errors = validateData(user)
        if (errors.length > 0) {
            throw {
                message: 'Please fill in the information completely.',
                errors: errors
            }
        }

        const results = await conn.query('INSERT INTO users SET ?', user)
        res.json({
            message: 'insert ok',
            data: results[0] 
        })
    } catch (error) {
        const errorMessage = error.message || 'something wrong'
        const errors = error.errors || []
        console.error('error message', error.message)
        res.status(500).json({
            message: errorMessage,
            errors: errors
        })
    }
})

// path =  GET /users/:id สำหรับการดึง users รายคนออกมา
app.get('/users/:id', async (req, res) => {
    try {
        let id = req.params.id
        const results = await conn.query('SELECT * FROM users WHERE id = ?', id)

        if (results[0].length == 0) {
            throw { statusCode: 404, message:'Data not Found' }   
        }

        res.json(results[0][0])
    } catch (error) {
        console.error('error message', error.message)
        let statusCode = error.statusCode || 500
        res.status(statusCode).json({
            message: 'something wrong',
            errorMessage: error.message
        })
    }
})

// path = PUT /users/:id
app.put('/users/:id', async (req, res) => {
    try {
        let id = req.params.id
        let updateUser = req.body

        const results = await conn.query(
            'UPDATE users SET ? WHERE id = ?', 
            [updateUser, id]
        )
        res.json({
            message: 'Update ok',
            data: results[0] 
        })
    } catch (error) {
        console.error('error message', error.message)
        res.status(500).json({
            message: 'something wrong'
        })
    }
})




// path = DELETE /users/:id
app.delete('/users/:id', async (req, res) => {
    try {
        let id = req.params.id
        const results = await conn.query('DELETE FROM users WHERE id = ?', id)
        res.json({
            message: 'Delete ok',
            data: results[0] 
        })
    } catch (error) {
        console.error('error message', error.message)
        res.status(500).json({
            message: 'something wrong'
        })
    }
})

app.listen(port, async (req, res) => {
    await initMySQL()
    console.log('Server is ranning on port' + port)
})