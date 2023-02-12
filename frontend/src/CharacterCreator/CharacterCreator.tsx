import './CharacterCreator.css';
// eslint-disable-next-line
import Sidebar from './Sidebar';

export default function CharacterCreator({charactername}: {charactername: string}) {
	return (
		<div className="background-main">
			<Sidebar>
				<div className="background-secondary shadow-2xl">
					<p className="titre text-8xl mt-11">fiche de personnage</p>
				</div>
			</Sidebar>
		</div>
	);
}
