export interface Api {
	get: (url: string) => Promise<any>;
}
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
	description: string | null
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
export interface Recipe {
	ing1: string
	ing2: string | null
	result: string
	ing1Quantity: number
	ing2Quantity: number | null
	resultQuantity: number
}
export interface ArmorInfo {
	name: string
	protection: number | null
	protectionType: string | null
	crafting: string | null
	effect: string | null
	id: number
	armor: number
	durability: number | null
	setId: number | null
}
export interface Shield {
	id: number
	name: string
	armor: number
}
export interface EnemyInfo {
	id: string;
	name: string;
}
export interface Enemy {
	id: string;
	difficulty: string;
	name: string;
	special: string;
	elemAttType?: string | null;
	horM?: string | null;
	fractureArmor?: string | null;
	fractureWeapon?: string | null;
	pierce?: string | null;
	cold?: string | null;
	fire?: string | null;
	health: number;
	armor: number;
	attack: number;
	elemAtt?: number | null;
	soulEater?: number | null;
	tgk?: number | null;
	instakill?: boolean | null;
}
