import SidebarHome from './SidebarCompte';
import Accordion from 'react-bootstrap/Accordion';
import Button from 'react-bootstrap/Button';
import characters from './backend.json';
import { useNavigate } from 'react-router';

export default function Compte({
	user,
	setModifying,
}: {
	user: string;
	setModifying: React.Dispatch<React.SetStateAction<string>>;
}) {
	const navigate = useNavigate();
	return (
		<div>
			<div>
				<div className="compte" style={{ backgroundColor: '#F5DEB3' }}>
					<SidebarHome>
						<div className="title_compte">
							<center>
								<b>Compte</b>
							</center>
						</div>
						<div className="Persona">
							<Accordion>
								{characters.map((element, index) => {
									if (user === element.user) {
										return (
											<Accordion.Item eventKey={`${index}`}>
												<Accordion.Header>Personnage: {element.name}</Accordion.Header>
												<Accordion.Body>
													profil: {element.profil_depart}, alignement: {element.alignement}, exp:{element.exp}
													<a href="/create" className="flex items-center p-2 space-x-3 rounded-md">
														<span
															className="button"
															onClick={() => {
																setModifying(element.name);
																navigate('/create');
															}}
														>
															Edit
														</span>
													</a>
												</Accordion.Body>
											</Accordion.Item>
										);
									}
									return <></>;
								})}
							</Accordion>
						</div>
					</SidebarHome>
				</div>
			</div>
		</div>
	);
}
