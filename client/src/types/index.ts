export interface Requirement {
	id: number
	name: string;
	quantity: number;
}

export interface Design {
	name: string
	id: number
	requirements: Requirement[]
}
