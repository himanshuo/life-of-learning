#pragma once
//backing store struct
typedef struct BackingStore
{
        int count;
        pthread_cond_t all_done;
        pthread_mutex_t count_lock;
};
void Barrier_Init(Barrier * b, int num_threads);
void Barrier_Wait(Barrier * b, int rank, int next);
// void Barrier_Destroy(Barrier * b);
// void Barrier_Reset(Barrier * b, int num_threads);
