import './CharacterCreator.css';
import React from 'react';
import Sidebar from './Sidebar';
import InputNumber from './InputNumber';
// import { readFileSync } from 'fs';

export default function CharacterCreator({characterName}: {characterName: string}) {
	// const [character, setCharacter] = React.useState(JSON.parse(readFileSync('./character.json', 'utf-8')));

	return (
		<div className="background-main">
			<Sidebar>
				<div className="background-secondary shadow-2xl">
					<p className="titre text-8xl mt-11 mb-6">fiche de personnage</p>
					<p className="titre text-3xl mt-1 mb-7">{characterName}</p>
					<div className="flex">
						<div className="columnLeft h-[100%]">
							<div className="flex h-[100%] m-1 rounded-lg bg-[#d1ad7d] shadow-[inset_0_-2px_4px_rgba(0,0,0,0.6)]">
								<div className="columnLeft m-2">
									<p className="mb-2">Experience</p>
									<p className="mb-2">Genre</p>
									<p className="mb-2">Race</p>
								</div>
								<div className="columnRight m-2">
									<InputNumber step={100} default={1000} />
									<br />
									<input className="mb-2" />
									<select className="selector" size={5}>
										<option value="1" selected={true}>
											Demi-elfe
										</option>
										<option value="2">Demi-orc</option>
										<option value="3">Elfe</option>
										<option value="4">Gnome</option>
										<option value="5">Halfelin</option>
										<option value="6">Humain</option>
										<option value="7">Nain</option>
										<option value="8">Aasimar</option>
										<option value="9">Changelin</option>
										<option value="10">Dhampir</option>
										<option value="11">Ifrit</option>
										<option value="12">Ondin</option>
										<option value="13">Oréade</option>
										<option value="14">Oréade</option>
										<option value="15">Sylphe</option>
										<option value="16">Tieffelin</option>
										<option value="17">Gnoll</option>
										<option value="18">Gobelin</option>
										<option value="19">Hobgobelin</option>
										<option value="20">Orc</option>
										<option value="21">Drow</option>
										<option value="22">Duergar</option>
										<option value="23">Svirfneblin</option>
										<option value="24">Félinidé</option>
										<option value="25">Homme-lézard</option>
										<option value="26">Grippli</option>
										<option value="27">Kitsune</option>
										<option value="28">Murinidé</option>
										<option value="29">Tengu</option>
										<option value="30">Troglodyte</option>
										<option value="31">Vanara</option>
										<option value="32">Kobold</option>
										<option value="33">Mutant draconique</option>
										<option value="34">Demi-dragon</option>
										<option value="35">Scutellaires</option>
										<option value="36">Réincarné</option>
									</select>
								</div>
							</div>
							<div className="flex h-[100%] m-1 rounded-lg bg-[#d1ad7d] shadow-[inset_0_-2px_4px_rgba(0,0,0,0.6)]">
								<div className="columnLeft m-2">
									<p className="mt-1">Origine géographique</p>
									<br className="mb-2" />
									<br className="mb-2" />
									<br className="mb-2" />
									<p className="mb-2">Origine sociale</p>
								</div>
								<div className="columnRight m-2">
									<select className="selector mb-2" size={4}>
										<option value="1" selected={true}>
											Arctique
										</option>
										<option value="2">Forêt</option>
										<option value="3">Désert</option>
										<option value="4">Plaine & savane</option>
										<option value="5">Marais</option>
										<option value="6">Montagnes & collines</option>
										<option value="7">Mer</option>
										<option value="8">Jungle</option>
										<option value="9">Outreterre</option>
										<option value="10">Urbain</option>
									</select>
									<select className="selector" size={4}>
										<option value="1" selected={true}>
											Paysan
										</option>
										<option value="2">Bourgeoisie</option>
										<option value="3">Aristocratie</option>
										<option value="4">Royautée</option>
									</select>
								</div>
							</div>
						</div>
						<div className="flex w-[50%] h-[100%]">
							<div className="columnLeft m-1 rounded-lg bg-[#d1ad7d] shadow-[inset_0_-2px_4px_rgba(0,0,0,0.6)]">
								<div className="flex">
									<div className="columnLeft mt-2 ml-2 mr-2">
										<p className="mb-2">Combat</p>
										<p className="mb-2">Connaissances</p>
										<p className="mb-2">Discrétion</p>
										<p className="mb-2">Endurance</p>
										<p className="mb-2">Force</p>
										<p className="mb-2">Habilité</p>
										<p className="mb-2">Magie</p>
										<p className="mb-2">Mouvement</p>
										<p className="mb-2">Perception</p>
										<p className="mb-2">Sociabilité</p>
										<p className="mb-2">Survie</p>
										<p className="mb-2">Tir</p>
										<p className="mb-2">Volonté</p>
									</div>
									<div className="columnRight mt-2 ml-2 mr-2">
										{/* <InputNumber step={1} default={character.competences.combat.score} /> */}
										{/* <InputNumber step={1} default={character.competences.connaissance.score} /> */}
										{/* <InputNumber step={1} default={character.competences.discretion.score} /> */}
										{/* <InputNumber step={1} default={character.competences.endurance.score} /> */}
										{/* <InputNumber step={1} default={character.competences.force.score} /> */}
										{/* <InputNumber step={1} default={character.competences.habilite.score} /> */}
										{/* <InputNumber step={1} default={character.competences.magie.score} /> */}
										{/* <InputNumber step={1} default={character.competences.mouvement.score} /> */}
										{/* <InputNumber step={1} default={character.competences.perception.score} /> */}
										{/* <InputNumber step={1} default={character.competences.sociabilite.score} /> */}
										{/* <InputNumber step={1} default={character.competences.survie.score} /> */}
										{/* <InputNumber step={1} default={character.competences.tir.score} /> */}
										{/* <InputNumber step={1} default={character.competences.volonté.score} /> */}
										<InputNumber step={1} default={0} />
										<InputNumber step={1} default={0} />
										<InputNumber step={1} default={0} />
										<InputNumber step={1} default={0} />
										<InputNumber step={1} default={0} />
										<InputNumber step={1} default={0} />
										<InputNumber step={1} default={0} />
										<InputNumber step={1} default={0} />
										<InputNumber step={1} default={0} />
										<InputNumber step={1} default={0} />
										<InputNumber step={1} default={0} />
										<InputNumber step={1} default={0} />
										<InputNumber step={1} default={0} />
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</Sidebar>
		</div>
	);
}
