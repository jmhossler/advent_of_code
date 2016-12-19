#include <algorithm>
#include <iostream>

int magic_missle(int turn, int hp, int mana, int boss_hp, int shield_timer,
                 int poison_timer, int recharge_timer);
int drain(int turn, int hp, int mana, int boss_hp, int shield_timer,
          int poison_timer, int recharge_timer);
int shield(int turn, int hp, int mana, int boss_hp, int shield_timer,
           int poison_timer, int recharge_timer);
int poison(int turn, int hp, int mana, int boss_hp, int shield_timer,
           int poison_timer, int recharge_timer);
int recharge(int turn, int hp, int mana, int boss_hp, int shield_timer,
             int poison_timer, int recharge_timer);


void apply_effects(int &hp, int &mana, int &boss_hp, int &shield_timer,
                   int &poison_timer, int &recharge_timer)
{
  shield_timer=std::max(shield_timer-1,0);

  if(poison_timer>0)
    {
      const int poison_damage=3;
      boss_hp-=poison_damage;
      --poison_timer;
    }

  if(recharge_timer>0)
    {
      const int recharge_mana=101;
      mana+=recharge_mana;
      --recharge_timer;
    }
}


int mana_costs(int turn, int hp, int mana, int boss_hp, int shield_timer,
               int poison_timer, int recharge_timer)
{
  --hp;
  if(hp<=0)
    return std::numeric_limits<int>::max();

  int recharge_cost=recharge(turn,hp,mana,boss_hp,shield_timer,poison_timer,
                             recharge_timer);
  int poison_cost=poison(turn,hp,mana,boss_hp,shield_timer,poison_timer,
                         recharge_timer);
  int shield_cost=shield(turn,hp,mana,boss_hp,shield_timer,poison_timer,
                         recharge_timer);
  int drain_cost=drain(turn,hp,mana,boss_hp,shield_timer,poison_timer,
                       recharge_timer);
  int magic_missle_cost=magic_missle(turn,hp,mana,boss_hp,shield_timer,
                                     poison_timer,recharge_timer);

  return
    std::min(recharge_cost,std::min(poison_cost,
                                        std::min(shield_cost,
                                                 std::min(drain_cost,
                                                          magic_missle_cost))));
}


int boss_turn(int turn, int mana_cost, int &hp, int &mana, int &boss_hp,
              int &shield_timer, int &poison_timer, int &recharge_timer)
{
  mana-=mana_cost;
  if (boss_hp<=0)
    return mana_cost;

  apply_effects(hp,mana,boss_hp,shield_timer,poison_timer,recharge_timer);

  if(boss_hp<=0)
    return mana_cost;
  const int boss_damage(9), shield_armor(7);
  // const int boss_damage(8), shield_armor(7);
  if(shield_timer>0)
    hp-=boss_damage-shield_armor;
  else
    hp-=boss_damage;

  if (hp<=0)
    return std::numeric_limits<int>::max();
  if (boss_hp<=0)
    return mana_cost;

  int result=mana_costs(turn+1,hp,mana,boss_hp,shield_timer,poison_timer,
                        recharge_timer);
  if (result==std::numeric_limits<int>::max())
    return result;
  return result + mana_cost;
}

int magic_missle(int turn, int hp, int mana, int boss_hp, int shield_timer,
                 int poison_timer, int recharge_timer)
{
  const int mana_cost=53;
  apply_effects(hp,mana,boss_hp,shield_timer,poison_timer,recharge_timer);

  if(mana<mana_cost)
    return std::numeric_limits<int>::max();

  boss_hp-=4;

  return boss_turn(turn, mana_cost, hp, mana, boss_hp, shield_timer,
                   poison_timer, recharge_timer);
}

int drain(int turn, int hp, int mana, int boss_hp, int shield_timer,
          int poison_timer, int recharge_timer)
{
  const int mana_cost=73;
  apply_effects(hp,mana,boss_hp,shield_timer,poison_timer,recharge_timer);

  if(mana<mana_cost)
    return std::numeric_limits<int>::max();

  boss_hp-=2;
  hp+=2;

  return boss_turn(turn, mana_cost, hp, mana, boss_hp, shield_timer,
                   poison_timer, recharge_timer);
}


int shield(int turn, int hp, int mana, int boss_hp, int shield_timer,
           int poison_timer, int recharge_timer)
{
  const int mana_cost=113;
  apply_effects(hp,mana,boss_hp,shield_timer,poison_timer,recharge_timer);

  if(mana<mana_cost || shield_timer!=0)
    return std::numeric_limits<int>::max();

  shield_timer=6;
  return boss_turn(turn, mana_cost, hp, mana, boss_hp, shield_timer,
                   poison_timer, recharge_timer);
}

int poison(int turn, int hp, int mana, int boss_hp, int shield_timer,
           int poison_timer, int recharge_timer)
{
  const int mana_cost=173;
  apply_effects(hp,mana,boss_hp,shield_timer,poison_timer,recharge_timer);

  if(mana<mana_cost || poison_timer!=0)
    return std::numeric_limits<int>::max();

  poison_timer=6;
  return boss_turn(turn, mana_cost, hp, mana, boss_hp, shield_timer,
                   poison_timer, recharge_timer);
}

int recharge(int turn, int hp, int mana, int boss_hp, int shield_timer,
             int poison_timer, int recharge_timer)
{
  const int mana_cost=229;
  apply_effects(hp,mana,boss_hp,shield_timer,poison_timer,recharge_timer);

  if(mana<mana_cost || recharge_timer!=0)
    return std::numeric_limits<int>::max();

  recharge_timer=5;
  return boss_turn(turn, mana_cost, hp, mana, boss_hp, shield_timer,
                   poison_timer, recharge_timer);
}



int main()
{
  std::cout << "result: " << mana_costs(0,50,500,51,0,0,0) << "\n";
}
