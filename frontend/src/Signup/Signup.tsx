import React, { useState } from 'react'
import Email from '../Login/Email'
import Password from '../Login/Password'
import CreateCompte from './CreateCompte'
import Title from './Title'
import ConfirmationPassword from './ConfirmationPassword'
import Name from './Name'
import './signup.css'

export default function SignupUniseed() {
    const [password, setPassword] = useState('')
    const [confirmation_password, setConfirmationPassword] = useState('')
    const [email, setEmail] = useState('')
    const [name, setName] = useState('')

    return (
        <div>
            <div className='split' style={{left:0, backgroundColor:"#8b4513"}}>
                <img src="rabbit.jpg" style={{height:920, width:600}}/>
            </div>
            <div className="split" style={{ right: 0, backgroundColor: "#e3b270"}}>
            <img src="dragon.png" style={{position:"absolute", left:420, width:80 }}/>
                <div className="centered">
                    <div className="mt-3 text-center">
                        <Title />
                    </div>
                    <div className="mt-4">
                        <Name name={name} onChange={setName} />
                    </div>
                    <div className="mt-4">
                        <Email email={email} onChange={setEmail} />
                    </div>
                    <div className="mt-4">
                        <Password password={password} onChange={setPassword} />
                    </div>
                    <div className="mt-4">
                        <ConfirmationPassword
                            password={confirmation_password}
                            onChange={setConfirmationPassword}
                        />
                    </div>
                    <div className="mt-4">
                        <CreateCompte
                            name={name}
                            email={email}
                            password={password}
                            confirmation_password={confirmation_password}
                        />
                    </div>
                </div>
            </div>
        </div>
    )
}
