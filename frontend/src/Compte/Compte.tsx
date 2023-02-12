import SidebarHome from './SidebarCompte';
import Accordion from 'react-bootstrap/Accordion';
import Button from 'react-bootstrap/Button';

export default function Compte() {

    return (
        <div>
            <div>
                <div className='compte' style={{ backgroundColor: "#F5DEB3" }}>
                    <SidebarHome>
                        <div className="title_compte">
                            <center><b>Compte</b></center>
                        </div>
                        <div className="Persona">
                            <Accordion>
                                <Accordion.Item eventKey="0">
                                    <Accordion.Header>Personnage: Darouine</Accordion.Header>
                                    <Accordion.Body>
                                        Darouine is a goblin. Rather, he has a strong character. He masters the art of combat with a two-handed weapon. He is equipped with an axe.
                                        <a href="/create" className="flex items-center p-2 space-x-3 rounded-md">
                                            <span className="button">Edit</span>
                                        </a>
                                    </Accordion.Body>
                                </Accordion.Item>
                                <Accordion.Item eventKey="1">
                                    <Accordion.Header>Personnage: Eclopia</Accordion.Header>
                                    <Accordion.Body>
                                        Eclopia is dark elf. He has a rather solitary character. He wields daggers with precision and controls fire magic.
                                        <a href="/create" className="flex items-center p-2 space-x-3 rounded-md">
                                            <span className="button">Edit</span>
                                        </a>
                                    </Accordion.Body>
                                </Accordion.Item>
                            </Accordion>
                        </div>
                    </SidebarHome>
                </div>
            </div>
        </div>
    );
}