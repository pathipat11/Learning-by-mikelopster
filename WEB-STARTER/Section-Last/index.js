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


const submitData = async () => {
    let firstNameDOM = document.querySelector('input[name=firstname]')
    let lastNameDOM = document.querySelector('input[name=lastname]')
    let ageDOM = document.querySelector('input[name=age]')

    let genderDOM = document.querySelector('input[name=gender]:checked') || {}
    let interestDOMs = document.querySelectorAll('input[name=interest]:checked') || {}

    let descriptionDOM = document.querySelector('textarea[name=description]')

    let messageDOM = document.getElementById('message')
    
    try {
        let interest = ''

        for (let i = 0; i < interestDOMs.length; i++) {
            interest += interestDOMs[i].value
            if (i != interestDOMs.length -1) {
                interest += ', '
            }
        }

        console.log('test')

        let userData = {
            firstname: firstNameDOM.value,
            lastname: lastNameDOM.value,
            age: ageDOM.value,
            gender: genderDOM.value,
            description: descriptionDOM.value,
            interests: interest
        }

        console.log('Submit Data', userData)
        
        const errors = validateData(userData)

        // if (errors.length > 0) {
        //     throw {
        //         message: 'Please fill in the information completely.',
        //         errors: errors
        //     }
        // }
        
        const response = await axios.post('http://localhost:8000/users', userData)
        console.log('response', response.data)

        messageDOM.innerText = 'Data saved successfully'
        messageDOM.className = 'message success'

    } catch (error) {
        console.log('error message', error.message)
        console.log('error', error.errors)
        if (error.response) {
            console.log(error.response)
            error.message = error.response.data.message
            error.errors = error.response.data.errors
        }


        let htmlData = '<div>'
        htmlData += `<div>${error.message}</div>`
        htmlData += '<ul>'
        for (let i = 0; i < error.errors.length; i++) {
            htmlData += `<li>${error.errors[i]}</li>`
        }
        htmlData += '</ul>'
        htmlData += '</div>'
        
        messageDOM.innerHTML = htmlData
        messageDOM.className = 'message danger'
    }
}

