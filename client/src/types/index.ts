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
export interface Trade {
	giveName: string
	getName: string
	giveQuantity: number
	getQuantity: number
	giveId: number
	getId: number
}
export interface MaterialInfo {
	name: string
	designs: DesignRequirement[]
}
export interface DesignRequirement {
	design: string
	designId: number
	quantity: number
}
