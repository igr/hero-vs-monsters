export type PlayerAttributes = {
  attackDamage: number,
  health: number,
  speed: number,
  name: string,
};

export type MonsterAttributes = PlayerAttributes & {
  speedDamage: number;
  cloneable: boolean;
};
