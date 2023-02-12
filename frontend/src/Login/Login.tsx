import React, { useState } from 'react';
import Email from './Email';
import Password from './Password';
import RememberMe from './RememberMe';
import Connection from './Connection';
import Title from './Title';
import Signup from './Signup';
import { useNavigate } from 'react-router';

export default function LoginUniseed({user, setUser} : {user: string, setUser: React.Dispatch<React.SetStateAction<string>>}) {
	const [password, setPassword] = useState('');
	const [email, setEmail] = useState('');
	const navigate = useNavigate();
	if (user === "") {
		navigate('/')
	}
	return (
		<div>
			<div className="split" style={{ left: 0, backgroundColor: '#8b4513' }}>
				<img className="rabbit" src="rabbit.jpg" alt="rabbit background" />
			</div>
			<div className="split" style={{ right: 0, backgroundColor: '#e3b270' }}>
				<img className="dragon_login" src="dragon.png" alt="dragon logo" />
				<div className="centered">
					<div className="text-center">
						<Title />
					</div>
					<div className="mt-3">
						<Email email={email} onChange={setEmail} />
					</div>
					<div className="mt-3">
						<Password password={password} onChange={setPassword} />
					</div>
					<div className="mt-3">
						<RememberMe />
					</div>
					<div className="mt-3">
						<Connection email={email} password={password} setUser={setUser}/>
					</div>
					<div
						className="mt-3"
						style={{
							position: 'absolute',
							top: 500,
							left: 100,
						}}
					>
						<Signup />
					</div>
				</div>
			</div>
		</div>
	);
}
