import React, { useState } from 'react';
import Email from '../Login/Email';
import Password from '../Login/Password';
import CreateCompte from './CreateCompte';
import Title from './Title';
import ConfirmationPassword from './ConfirmationPassword';
import Name from './Name';
import './signup.css';
import { Navigate, useNavigate } from 'react-router';

export default function SignupUniseed({user} : {user: string}) {
	const [password, setPassword] = useState('');
	const [confirmationPassword, setConfirmationPassword] = useState('');
	const [email, setEmail] = useState('');
	const [name, setName] = useState('');
	const navigate = useNavigate();
	if (user === "") {
		navigate('/compte')
	}
	return (
		<div>
			<div className="split" style={{ left: 0, backgroundColor: '#8b4513' }}>
				<img className="rabbit" src="rabbit.jpg" alt="rabbit background" />
			</div>
			<div className="split" style={{ right: 0, backgroundColor: '#e3b270' }}>
				<img className="dragon" src="dragon.png" alt="dragon logo" />
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
						<ConfirmationPassword password={confirmationPassword} onChange={setConfirmationPassword} />
					</div>
					<div className="mt-4">
						<CreateCompte name={name} email={email} password={password} confirmationPassword={confirmationPassword} />
					</div>
				</div>
			</div>
		</div>
	);
}
