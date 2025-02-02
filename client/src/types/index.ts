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
export interface WeaponInfo {
	attackSpeed: number
	name: string
	damage: number
	s1: number | null
	s2: number | null
	s3: number | null
	s4: number | null
	s5: number | null
}
