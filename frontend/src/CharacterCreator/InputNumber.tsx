interface Props {
    step: number
    default: number
}

export default function InputNumber(props: Props) {
	return (
		<input
			type="number"
			min="0"
			step={props.step}
            defaultValue={props.default}
            className="mb-2"
			onKeyDown={(evt) => {
				(evt.key < '0' || evt.key > '9') &&
					evt.key !== 'Backspace' &&
					evt.key !== 'ArrowRight' &&
					evt.key !== 'ArrowLeft' &&
					evt.key !== 'Delete' &&
					evt.preventDefault();
			}}
		/>
	);
}
