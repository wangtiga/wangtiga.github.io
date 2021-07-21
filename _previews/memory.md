---
layout: post
title:  "[译]What a C programmer should know about memory"
date:   2020-04-29 12:00:00 +0800
tags:   todo
---

* category
{:toc}


# What a C programmer should know about memory [^vavrusamemory] [^v2exmemory]

![](https://upload.wikimedia.org/wikipedia/commons/thumb/8/8c/Oriental_weapons.jpg/800px-Oriental_weapons.jpg)  
Source: Weapons by  [T4LLBERG](https://www.flickr.com/photos/t4llberg/5828882565), on Flickr (CC-BY-SA)

In 2007, Ulrich Drepper wrote a  _“[What every programmer should know about memory](http://www.akkadia.org/drepper/cpumemory.pdf)“_. Yes, it’s a wee long-winded, but it’s worth its salt. Many years and “every programmer should know about” articles later, the concept of virtual memory is still elusive to many, as if it was  _a kind of magic_. Awww, I couldn’t resist the reference. Even the validity of the original article was  [questioned](https://stackoverflow.com/questions/8126311/what-every-programmer-should-know-about-memory)  many years later.  _What gives?_

> “North bridge? What is this crap? That ain’t street-fighting.”

I’ll try to convey the practical side of things (i.e. what can you do) from  _“getting your fundamentals on a lock”_, to more fun stuff. Think of it as a glue between the original article, and the things that you use every day. The examples are going to be C99 on Linux, but a lot of topics are universal.  _EDIT: I don’t have much knowledge about Windows, but I’d be thrilled to link an article which explains it. I tried my best to mention which functions are platform-specific, but again I’m only a human. If you find a discrepancy, please let me know._

Without further ado, grab a cup of coffee and let’s get to it.

## Understanding virtual memory - the plot thickens

Unless you’re dealing with some embedded systems or kernel-space code, you’re going to be working in protected mode. This is  _awesome_, since your program is guaranteed to have it’s own [virtual] address space. The word  _“virtual”_  is important here. This means, among other things, that you’re not bounded by the available memory, but also not entitled to any. In order to use this space, you have to ask the OS to back it with something real, this is called  _mapping_. A backing can be either a physical memory (not necessarily RAM), or a persistent storage. The former is also called an  _“anonymous mapping”_. But hold your horses.

The virtual memory allocator (VMA)  _may_  give you a memory it doesn’t have, all in a vain hope that you’re not going to use it. Just like banks today. This is called  _[overcommiting](https://www.kernel.org/doc/Documentation/vm/overcommit-accounting)_, and while it has legitimate applications  _(sparse arrays)_, it also means that the memory allocation is not going to simply say  **“NO”**.

```
char *block = malloc(1024 * sizeof(char));
if (block == NULL) {
	return -ENOMEM; /* Sad :( */
}

```

The  `NULL`  return value checking is a good practice, but it’s not as powerful as it once was. With the overcommit, the OS may give your memory allocator a valid pointer to memory, but if you’re going to access it - dang*. The  _dang_  in this case is platform-specific, but generally an  [OOM killer](http://www.win.tue.nl/~aeb/linux/lk/lk-9.html#ss9.6)  killing your process.

*  —  _This is an oversimplification, as timbatron  [noted](https://www.reddit.com/r/C_Programming/comments/2ya9gl/what_a_c_programmer_should_know_about_memory/cp8mpau), and it’s further explained in the  [“Demand paging explained”](https://marek.vavrusa.com/memory/#pagefault)  section. But I’d like to go through the well-known stuff first before we delve into specifics._

### Detour - a process memory layout

The layout of a process memory is well covered in the  _[Anatomy of a Program in Memory](http://duartes.org/gustavo/blog/post/anatomy-of-a-program-in-memory/)_  by Gustavo Duarte, so I’m going to quote and reference to the original article, I hope it’s a  _fair use_. I have only a few minor quibbles, for one it covers only a x86-32 memory layout, but fortunately nothing much has changed for x86-64. Except that a process can use much more space — the whopping 48 bits on Linux.

  
[![x86-32 Linux address space layout, source: duartes.org/gustavo/blog/post/anatomy-of-a-program-in-memory](http://static.duartes.org/img/blogPosts/linuxFlexibleAddressSpaceLayout.png)](http://duartes.org/gustavo/blog/post/anatomy-of-a-program-in-memory/)  
Source: Linux address space layout by  [Gustavo Duarte](http://duartes.org/gustavo/blog/post/anatomy-of-a-program-in-memory/)

It also shows the memory mapping segment (MMS) growing down, but that may not always be the case. The MMS usually starts ([x86/mm/mmap.c:113](http://lxr.free-electrons.com/source/mm/mmap.c#L1953)  and  [arch/mm/mmap.c:1953](http://lxr.free-electrons.com/source/arch/x86/mm/mmap.c#L113)) at a randomized address just below the lowest address of the stack.  _Usually_, because it may start above the stack and grow upwards  _iff_  the stack limit is large (or  _unlimited_), or the compatibility layout is enabled.  _How is this important?_  It’s not, but it helps to give you an idea about the  [free address ranges](https://marek.vavrusa.com/memory/#mmap-fun).

Looking at the diagram, you can see three possible variable placements: the process data segment  _(static storage or heap allocation)_, the memory mapping segment, and the stack. Let’s start with that one.

## Understanding stack allocation

_Utility belt:_

-   [`alloca()`  - allocate memory in the stack frame of the caller](https://linux.die.net/man/3/alloca)
-   [`getrlimit()`  - get/set resource limits](https://linux.die.net/man/2/getrlimit)
-   [`sigaltstack()`  - set and/or get signal stack context](https://linux.die.net/man/2/sigaltstack)

The stack is kind of easy to digest, everybody knows how to make a variable on the stack right? Here are two:

```
int stairway = 2;
int heaven[] = { 6, 5, 4 };

```

The validity of the variables is limited by scope. In C, that means this:  `{}`. So each time a closing curly bracket comes, a variable dies. And then there’s  [`alloca()`](https://linux.die.net/man/3/alloca), which allocates memory dynamically in the current  _stack frame_. A stack frame is not (entirely) the same thing as memory frame (aka  _physical_  page), it’s simply a group of data that gets pushed onto the stack (function, parameters, variables…). Since we’re on the top of the stack, we can use the remaining memory up to the stack size limit.

This is how variable-length arrays (VLA), and also  [`alloca()`](https://linux.die.net/man/3/alloca)  work, with one difference -  _VLA_  validity is limited by the scope, alloca’d memory persists until the current function returns (or  _unwinds_  if you’re feeling sophisticated). This is no language lawyering, but a real issue if you’re using the alloca inside a loop, as you don’t have any means to free it.

```
void laugh(void) {
	for (unsigned i = 0; i < megatron; ++i) {
	    char *res = alloca(2);
	    memcpy(res, "ha", 2);
	    char vla[2] = {'h','a'}
	} /* vla dies, res lives */
} /* all allocas die */

```

Neither VLA or alloca play nice with large allocations, because you have almost no control over the available stack memory and the allocation past the stack limits leads to the jolly stack overflow. There are two ways around it, but neither is practical.

The first idea is to use a  [`sigaltstack()`](https://linux.die.net/man/2/sigaltstack)  to catch and handle the  `SIGSEGV`. However this just lets you catch the stack overflow.

The other way is to compile with  _split-stacks_. It’s called this way, because it really splits the monolithic stack into a linked-list of smaller stacks called  _stacklets_. As far as I know,  [GCC](https://gcc.gnu.org/wiki/SplitStacks)  and  [clang](https://llvm.org/releases/3.0/docs/SegmentedStacks.html)  support it with the  [`-fsplit-stack`](https://gcc.gnu.org/wiki/SplitStacks)  option. In theory this also improves memory consumption and reduces the cost of creating threads — because the stack can start really small and grow on demand. In reality, expect compatibility issues, as it needs a split-stack aware linker (i.e. gold) to play nice with the split-stack  _unaware_  libraries, and performance issues (the  _“hot split”_  problem in Go is  [nicely explained](http://agis.io/2014/03/25/contiguous-stacks-in-go.html)  by Agis Anastasopoulos).

## Understanding heap allocation

_Utility belt:_

-   [`brk(), sbrk()`  - manipulate the data segment size](https://linux.die.net/man/2/sbrk)
-   [`malloc()`  family - portable libc memory allocator](https://linux.die.net/man/3/malloc)

The heap allocation can be as simple as moving a  [program break](https://linux.die.net/man/2/sbrk)  and claiming the memory between the old position, and the new position. Up to this point, a heap allocation is as fast as stack allocation  _(sans the paging, presuming the stack is already locked in memory)_. But there’s a  cat, I mean catch, dammit.

```
char *block = sbrk(1024 * sizeof(char));

```

⑴ we can’t reclaim unused memory blocks, ⑵ is not thread-safe since the heap is shared between threads, ⑶ the interface is hardly portable, libraries must not touch the break

> `man 3 sbrk`  — Various systems use various types for the argument of sbrk(). Common are int, ssize_t, ptrdiff_t, intptr_t.

For these reasons libc implements a centralized interface for memory allocation. The  [implementation varies](https://en.wikibooks.org/wiki/C_Programming/C_Reference/stdlib.h/malloc#Implementations), but it provides you a thread safe memory allocation of any size … at a cost. The cost is latency, as there is now locking involved, data structures keeping the information about used / free blocks and an extra memory overhead. The heap is not used exclusively either, as the memory mapping segment is often utilised for large blocks as well.

> `man 3 malloc`  — Normally,  `malloc()`  allocates memory from the heap, … when allocating blocks of memory larger than MMAP_THRESHOLD, the glibc  `malloc()`  implementation allocates the memory as a private anonymous mapping.

As the heap is always contiguous from  `start_brk`  to  `brk`, you can’t exactly punch holes through it and reduce the data segment size. Imagine the following scenario:

```
char *truck = malloc(1024 * 1024 * sizeof(char));
char *bike  = malloc(sizeof(char));
free(truck);

```

The heap [allocator] moves the  _brk_  to make space for  `truck`. The same for the  `bike`. But after the  `truck`  is freed, the  _brk_  can’t be moved down, as it’s the  `bike`  that occupies the highest address. The result is that your process  _can_  reuse the former  `truck`  memory, but it can’t be returned to the system until the  `bike`  is freed. But presuming the  `truck`  was mmaped, it wouldn’t reside in the heap segment, and couldn’t affect the program break. Still, this trick doesn’t prevent the holes created by small allocations (in another words “cause  _fragmentation_”).

Note that the  `free()`  doesn’t always try to shrink the data segment, as that is a  [potentially expensive operation](https://marek.vavrusa.com/memory/#pagefault). This is a problem for long-running programs, such as daemons. A GNU extension, called  [`malloc_trim()`](https://linux.die.net/man/3/malloc_trim), exists for releasing memory from the top of the heap, but it can be painfully slow. It hurts  _real bad_  for a lot of small objects, so it should be used sparingly.

### When to bother with a custom allocator

There are a few practical use cases where a GP allocator falls short — for example an allocation of a  _large_  number of  _small_  fixed-size chunks. This might not look like a typical pattern, but it is very frequent. For example, lookup data structures like trees and tries typically require nodes to build hierarchy. In this case, not only the fragmentation is the problem, but also the data locality. A cache-efficient data structure packs the keys together (preferably on the same page), instead of mixing it with data. With the default allocator, there is  _no guarantee_  about the locality of the blocks from the subsequent allocations. Even worse is the space overhead for allocating small units. Here comes the solution!

[![](https://farm4.staticflickr.com/3090/2642284820_e3bd840bca_z.jpg?zz=1)](https://flic.kr/p/52updQ)  
Source: Slab by  [wadem](https://www.flickr.com/people/99174151@N00/), on Flickr (CC-BY-SA)

#### Slab allocator

_Utility belt:_

-   [`posix_memalign()`  - allocate aligned memory](https://linux.die.net/man/3/posix_memalign)

The principle of slab allocation was described by  [Bonwick](https://www.usenix.org/legacy/publications/library/proceedings/bos94/full_papers/bonwick.a)  for a kernel object cache, but it applies for the user-space as well.  _Oh-kay_, we’re not interested in pinning slabs to CPUs, but back to the gist — you ask the allocator for a  _slab_  of memory, let’s say a whole  _page_, and you cut it into many fixed-size pieces. Presuming each piece can hold at least a pointer or an integer, you can link them into a list, where the  _list head_  points to the  _first free_  element.

```
/* Super-simple slab. */
struct slab {
	void **head;
};

/* Create page-aligned slab */
struct slab *slab = NULL;
posix_memalign(&slab, page_size, page_size);
slab->head = (void **)((char*)slab + sizeof(struct slab));

/* Create a NULL-terminated slab freelist */
char* item = (char*)slab->head;
for(unsigned i = 0; i < item_count; ++i) {
	*((void**)item) = item + item_size;
	item += item_size;
}
*((void**)item) = NULL;

```

Allocation is then as simple as popping a list head. Freeing is equal to as pushing a new list head. There is also a neat trick. If the slab is aligned to the  `page_size`  boundary, you can get the slab pointer as cheaply as  [rounding down](https://stackoverflow.com/a/2601527/4591872)  to the  `page_size`.

```
/* Free an element */
struct slab *slab = (void *)((size_t)ptr & PAGESIZE_BITS);
*((void**)ptr) = (void*)slab->head;
slab->head = (void**)ptr;

/* Allocate an element */
if((item = slab->head)) {
	slab->head = (void**)*item;
} else {
	/* No elements left. */
}

```

Great, but what about binning, variable size storage, cache aliasing and caffeine, …? Peek at  [my old implementation](https://github.com/CZNIC-Labs/knot/blob/1.5/src/common-knot/slab/slab.h)  for  [Knot DNS](https://github.com/CZNIC-Labs/knot)  to get the idea, or use a library that implements it. For example, *gasp*, the glib implementation has a  [tidy documentation](https://developer.gnome.org/glib/stable/glib-Memory-Slices.html)  and calls it  _“memory slices”_.

#### Memory pools

_Utility belt:_

-   [`obstack_alloc()`  - allocate memory from object stack](https://www.gnu.org/software/libc/manual/html_node/Obstacks.html)

As with the slab, you’re going to outsmart the GP allocator by asking it for whole chunks of memory only. Then you just slice the cake until it runs out, and then ask for a new one. And another one. When you’re done with the cakes, you call it a day and free everything in one go.

Does it sound obvious and stupid simple? Because it is, but that’s what makes it great for specific use cases. You don’t have to worry about synchronisation, not about freeing either. There are no use-after-free bugs, data locality is much more predictable, there is  _almost zero_  overhead for small fragments.

The pattern is surprisingly suitable for many tasks, ranging from short-lived repetitive (i.e.  _“network request processing”_), to long-lived immutable data (i.e.  _“frozen set”_). You don’t have to free  _everything_  either. If you can make an educated guess on how much memory is needed on average, you can just free the excess memory and reuse. This reduces the memory allocation problem to simple pointer arithmetic.

And you’re in luck here, as the GNU libc provides, *whoa*, an actual API for this. It’s called  [_obstacks_](https://www.gnu.org/software/libc/manual/html_node/Obstacks.html), as in “stack of objects”. The HTML documentation formatting is a bit underwhelming, but minor quibbles aside — it allows you to do both pool allocation, and full or partial unwinding.

```
/* Define block allocator. */
#define obstack_chunk_alloc malloc
#define obstack_chunk_free free

/* Initialize obstack and allocate a bunch of animals. */
struct obstack animal_stack;
obstack_init (&animal_stack);
char *bob = obstack_alloc(&animal_stack, sizeof(animal));
char *fred = obstack_alloc(&animal_stack, sizeof(animal));
char *roger = obstack_alloc(&animal_stack, sizeof(animal));

/* Free everything after fred (i.e. fred and roger). */
obstack_free(&animal_stack, fred);

/* Free everything. */
obstack_free(&animal_stack, NULL);

```

There is one more trick to it, you can grow the object on the top of the stack. Think buffering input, variable-length arrays, or just a way to combat the  `realloc()-strcpy()`  pattern.

```
/* This is wrong, I better cancel it. */
obstack_grow(&animal_stack, "long", 4);
obstack_grow(&animal_stack, "fred", 5);
obstack_free (&animal_stack, obstack_finish(&animal_stack));

/* This time for real. */
obstack_grow(&animal_stack, "long", 4);
obstack_grow(&animal_stack, "bob", 4);
char *result = obstack_finish(&animal_stack);
printf("%s\n", result); /* "longbob" */

```

#### Demand paging explained

_Utility belt:_

-   [`mlock()`  - lock/unlock memory](https://linux.die.net/man/2/mlock)
-   [`madvise()`  - give advice about use of memory](https://linux.die.net/man/2/madvise)

One of the reasons why the GP memory allocator doesn’t immediately return the memory to the system is, that it’s costly. The system has to do two things: ⑴ establish the mapping of a  _virtual_  page to  _real_  page, and ⑵ give you a blanked  _real_  page. The  _real_  page is called  _frame_, now you know the difference. Each frame must be sanitized, because you don’t want the operating system to leak your secrets to another process, would you? But here’s the trick, remember the  _overcommit_? The virtual memory allocator honours the only the first part of the deal, and then plays some  _“now you see me and now you don’t”_  shit — instead of pointing you to a real page, it points to a  _special_  page  `0`.

Each time you try to access the special page, a page  _fault_  occurs, which means that: the kernel pauses process execution and fetches a real page, then it updates the page tables, and resumes like nothing happened. That’s about the best explanation I could muster in one sentence,  [here’s](http://duartes.org/gustavo/blog/post/how-the-kernel-manages-your-memory/)  more detailed one. This is also called  _“demand paging”_  or  _“lazy loading”_.

> The Spock said that  _“one man cannot summon the future”_, but here you can pull the strings.

The memory manager is no oracle and it makes very conservative predictions about how you’re going to access memory, but you may know better. You can  [lock](https://linux.die.net/man/2/mlock)  the contiguous memory block in  _physical_  memory, avoiding further page faulting:

```
char *block = malloc(1024 * sizeof(char));
mlock(block, 1024 * sizeof(char));

```

*psst*, you can also give an  [advise](https://linux.die.net/man/2/madvise)  about your memory usage pattern:

```
char *block = malloc(1024 * sizeof(block));
madvise(block, 1024 * sizeof(block), MADV_SEQUENTIAL);

```

The interpretation of the actual advice is platform-specific, the system may even choose to ignore it altogether, but most of the platforms play nice. Not all advices are well-supported, and some even change semantics (`MADV_FREE`  drops dirty private memory), but the  `MADV_SEQUENTIAL`,  `MADV_WILLNEED`, and  `MADV_DONTNEED`  holy trinity is what you’re going to use most.

## Fun with  flags  memory mapping

_Utility belt:_

-   [`sysconf()`  - get configuration information at run time](https://linux.die.net/man/3/sysconf)
-   [`mmap()`  - map virtual memory](https://linux.die.net/man/2/mlock)
-   [`mincore()`  - determine whether pages are resident in memory](https://linux.die.net/man/2/mincore)
-   [`shmat()`  - shared memory operations](https://linux.die.net/man/2/shmat)

There are several things that the memory allocator  _just can’t_  do, memory maps to to rescue! To pick one, the fact that you can’t choose the allocated address range. For that we’re willing to sacrifice some comfort — we’re going to be working with whole pages from now on. Just to make things clear, a page is usually a 4K block, but you shouldn’t rely on it and use  [`sysconf()`](https://linux.die.net/man/3/sysconf)  to discover it.

```
long page_size = sysconf(_SC_PAGESIZE); /* Slice and dice. */

```

Side note — even if the platform advertises a uniform page size, it may not do so in the background. For example a Linux has a concept of  [transparent huge pages](https://lwn.net/Articles/423584/)  (THP) to reduce the cost of address translation and page faulting for contiguous blocks. This is however questionable, as the huge contiguous blocks become scarce when the physical memory gets fragmented. The cost of faulting a huge page also increases with the page size, so it’s not very efficient for “small random I/O” workload. This is unfortunately transparent to you, but there is a Linux-specific  `mmap`  option  [`MAP_HUGETLB`](https://www.kernel.org/doc/Documentation/vm/hugetlbpage.txt)  that allows you to use it explicitly, so you should be aware of the costs.

### Fixed memory mappings

Say you want to do fixed mapping for a poor man’s IPC for example, how do you choose an address? On x86-32 bit it’s a risky proposal, but on the 64-bit, an address around 2/3rds of the  `TASK_SIZE`  (highest usable address of the user space process) is a safe bet. You can get away without fixed mapping, but then forget pointers in your shared memory.

```
#define TASK_SIZE 0x800000000000
#define SHARED_BLOCK (void *)(2 * TASK_SIZE / 3)

void *shared_cats = shmat(shm_key, SHARED_BLOCK, 0);
if(shared_cats == (void *)-1) {
    perror("shmat"); /* Sad :( */
}


```

Okay, I get it, this is hardly a portable example, but you get the gist. Mapping a fixed address range is usually considered unsafe at least, as it doesn’t check whether there is something already mapped or not. There is a  [`mincore()`](https://linux.die.net/man/2/mincore)  function to tell you whether a page is mapped or not, but you’re out of luck in multi-threaded environment.

However, fixed-address mapping is not only useful for unused address ranges, but for the  _used_  address ranges as well. Remember how the memory allocator used  `mmap()`  for bigger chunks? This makes efficient sparse arrays possible thanks to the on-demand paging. Let’s say you have created a sparse array, and now you want to free some data, but how to do that? You can’t exactly  `free()`  it, and  `munmap()`  would render it unusable. You could use the  [`madvise()`](https://marek.vavrusa.com/memory/#pagefault)  `MADV_FREE / MADV_DONTNEED`  to mark the pages free, this is the best solution performance-wise as the pages don’t have to be faulted in, but the semantics of the advice  [differs](https://lwn.net/Articles/591214/)  is implementation-specific.

A portable approach is to map over the sucker.

```
void *array = mmap(NULL, length, PROT_READ|PROT_WRITE,
                   MAP_ANONYMOUS, -1, 0);

/* ... some magic gone awry ... */

/* Let's clear some pages. */
mmap(array + offset, length, MAP_FIXED|MAP_ANONYMOUS, -1, 0);

```

This is equivalent to unmapping the old pages and mapping them again to that  _special page_. How does this affect the perception of the process memory consumption — the process still uses the same amount of virtual memory, but the resident  _[in physical memory]_  size lowers. This is as close to memory  _hole punching_  as we can get.

### File-backed memory maps

_Utility belt:_

-   [`msync()`  - synchronize a file with memory map](https://linux.die.net/man/2/msync)
-   [`ftruncate()`  - truncate a file to a specified length](https://linux.die.net/man/2/ftruncate)
-   [`vmsplice()`  - splice user pages into a pipe](https://linux.die.net/man/2/vmsplice)

So far we’ve been all about anonymous memory, but it’s the file-backed memory mapping that really shines in the 64 bit address space, as it provides you with intelligent caching, synchronization and copy-on-write. Maybe that’s too much.

> To most people, LMDB is magic performance sprinkles compared to using the filesystem directly. ;)
> 
> —  [Baby_Food](https://www.reddit.com/r/programming/comments/2vyzer/what_every_programmer_should_know_about/comhq3s)  on r/programming

The file-backed shared memory maps add novel mode  `MAP_SHARED`, which means that the changes you make to the pages will be written back to the file, therefore shared with other processes. The decision of  _when_  to synchronize is left up to the memory manager, but fortunately there’s a  [`msync()`](https://linux.die.net/man/2/msync)  function to enforce the synchronization with the backing store. This is great for the databases, as it guarantees durability of the written data. But not everyone needs that, it’s perfectly okay  **not**  to sync if the durability isn’t required, you’re not going to lose write visibility. This is thanks to the page cache, and it’s good because you can use memory maps for cheap IPC for example.

```
/* Map the contents of a file into memory (shared). */
int fd = open(...);
void *db = mmap(NULL, file_size, PROT_READ|PROT_WRITE,
                MAP_SHARED, fd, 0);
if (db == (void *)-1) {
	/* Mapping failed */
}

/* Write to a page */
char *page = (char *)db;
strcpy(page, "bob");
/* This is going to be a durable page. */
msync(page, 4, MS_SYNC);
/* This is going to be a less durable page. */
page = page + PAGE_SIZE;
strcpy(page, "fred");
msync(page, 5, MS_ASYNC);

```

Note that you can’t map more bytes than the file actually has, so you can’t grow it or shrink it this way. You can however create (or grow) a sparse file in advance with  [`ftruncate()`](https://linux.die.net/man/2/ftruncate). The downside is, that it makes compaction harder, as the ability to punch holes through a sparse file depends both on the file system, and the platform.

The  [`fallocate(FALLOC_FL_PUNCH_HOLE)`](https://linux.die.net/man/2/fallocate)  on Linux is your best chance, but the most portable (and easiest) way is to make a copy of the file without the trimmed stuff.

```
/* Resize the file. */
int fd = open(...);
ftruncate(fd, expected_length);

```

Accessing a file memory map doesn’t exclude using it as a file either. This is useful for implementing a split access, where you map the file read only, but write to the file using a standard file API. This is good for security (as the exposed map is write-protected), but there’s more to it. The  [`msync()`](https://linux.die.net/man/2/msync)  implementation is not defined, so the  `MS_SYNC`  may very well be just a sequence of synchronous writes.  _Yuck._  In that case, it may be faster to use a regular file APIs to do an asynchronous  [`pwrite()`](https://linux.die.net/man/2/pwrite)  and  [`fsync() / fdatasync()`](https://linux.die.net/man/2/fsync)  for synchronisation or cache invalidation.

As always there is a caveat — the system has to have a  _unified buffer cache_. Historically, a page cache and block device cache  _(raw blocks)_  were two different things. This means that writing to a file using a standard API and reading it through a memory map is not coherent[1](https://marek.vavrusa.com/memory/#fn:coherent), unless you invalidate buffers after each write.  _Uh oh._  On the other hand, you’re in luck unless you’re running OpenBSD or Linux < 2.4.

#### Copy-on-write

So far this was about shared memory mapping. But you can use the memory mapping in another way — to map a shared copy of a file, and make modifications without modifying the backing store. Note that the pages are not duplicated immediately, that wouldn’t make sense, but in the moment you modify them. This is not only useful for forking processes or loading shared libraries, but also for working on a large set of data in-place, from multiple processes at once.

```
int fd = open(...);

/* Copy-on-write mapping */
void *db = mmap(NULL, file_size, PROT_READ|PROT_WRITE,
                    MAP_PRIVATE, fd, 0);
if (db == (void *)-1) {
	/* Mapping failed */
}

/* This page will be copied as soon as we write to it */
char *page = (char *)db;
strcpy(page, "bob");

```

#### Zero-copy streaming

Since the file is essentially a memory, you can stream it to pipes (that includes sockets), zero-copy style. Unlike the  [`splice()`](https://linux.die.net/man/2/splice), this plays well with the copy-on-write modification of the data.  _Disclaimer: This is for Linux folks only!_

```
int sock = get_client();
struct iovec iov = { .iov_base = cat_db, .iov_len = PAGE_SIZE };
int ret = vmsplice(sock, &iov, 1, 0);
if (ret != 0) {
	/* No streaming :( */
}

```

#### When mmap() isn’t the holy grail

There are pathological cases where mmapping a file is going to be much worse than the usual approach. A rule of thumb is that handling a page fault is slower than simply reading a file block, on the basis that it has to read the file block  _and_  do something more. In reality though, mmapped I/O may be faster as it avoids double or triple caching of the data, and does read-ahead in the background. But there are times when this is going to hurt. One such example is “small random reads in a file larger than available memory”. In this case the system reads ahead blocks that are likely not to be used, and each access is going to page fault instead. You can combat this to a degree with  [`madvise()`](https://linux.die.net/man/2/madvise).

Then there’s TLB thrashing. Translation of each virtual page to a frame is hardware-assisted, and the CPU keeps a cache of latest translations — this is the Translation Lookaside Buffer. A random access to a larger number of pages than the cache can hold inevitably leads to  _“thrashing”_, as the system has to do the translation by walking the page tables. For other cases, the solution is to use  [huge pages](https://wiki.debian.org/Hugepages), but it’s not going to cut it, as loading  _megabytes_  worth of data just to access a few odd bytes has even more detrimental effect.

## Understanding memory consumption

_Utility belt:_

-   [`vmtouch`  - portable virtual memory toucher](http://hoytech.com/vmtouch/)

The concept of shared memory renders the traditional approach — measuring resident size &mdash obsolete, as there’s no just quantification on the amount exclusive for your process. That leads to confusion and horror, which can be two-fold:

> With mmap’d I/O, our app now uses almost zero memory.  
> — CorporateGuy
> 
> Helpz! My process writing to shared map leaks so much memory!!1  
> — HeavyLifter666

There are two states of pages,  `clean`  and  `dirty`. The difference is that a dirty page has to be flushed to permanent storage before it can be reclaimed. The  `MADV_FREE`  advice uses this as a cheap way to free memory just by clearing the dirty bit instead of updating the page table entry. In addition, each page can be either  `private`  or  `shared`, and this is where things get confusing.

Both claims are [sort of] true, depending on the perspective. Do pages in the buffer cache count towards the process memory consumption? What about when a process dirties file-backed pages that end up in the buffer cache? How to make something out of this madness?

Imagine a process,  _the_eye_, writing to the shared map of mordor. Writing to the shared memory doesn’t count towards Rss, right?

```
$ ps -p $$ -o pid,rss
  PID   RSS
17906 1574944 # <-- WTF?

```

Err, back to the drawing board.

#### PSS (Proportional Set Size)

Proportional Set Size counts the private maps and adds a  _portion_  of the shared maps. This is as fair as we can get when talking about memory. By  _“portion”_, I mean a size of a shared map, divided by the number of processes sharing it. Let’s see an example, we have an application that does read/write to a shared memory map.

```
$ cat /proc/$$/maps
00400000-00410000         r-xp 0000 08:03 1442958 /tmp/the_eye
00bda000-01a3a000         rw-p 0000 00:00 0       [heap]
7efd09d68000-7f0509d68000 rw-s 0000 08:03 4065561 /tmp/mordor.map
7f0509f69000-7f050a108000 r-xp 0000 08:03 2490410 libc-2.19.so
7fffdc9df000-7fffdca00000 rw-p 0000 00:00 0       [stack]
... snip ...

```

Here’s the simplified breakdown of each map, the first column is the address range, the second is permissions information, where  `r`  stands for  _read_,  `w`  stands for  _write_,  `x`  means  _executable_  — so far the classic stuff —  `s`  is  _shared_  and  `p`  is  _private_. Then there’s offset, device, inode, and finally a  _pathname_.  [Here’s](https://www.kernel.org/doc/Documentation/filesystems/proc.txt)  the documentation,  _massively_  comprehensive.

I admit I’ve snipped the not-so-interesting bits from the output. Read  [FAQ (Why is “strict overcommit” a dumb idea?)](http://landley.net/writing/memory-faq.txt)  if you’re interested why the shared libraries are mapped as private, but it’s the map of mordor that interests us:

```
$ grep -A12 mordor.map /proc/$$/smaps
Size:           33554432 kB
Rss:             1557632 kB
Pss:             1557632 kB
Shared_Clean:          0 kB
Shared_Dirty:          0 kB
Private_Clean:   1557632 kB
Private_Dirty:         0 kB
Referenced:      1557632 kB
Anonymous:             0 kB
AnonHugePages:         0 kB
Swap:                  0 kB
KernelPageSize:        4 kB
MMUPageSize:           4 kB
Locked:                0 kB
VmFlags: rd wr sh mr mw me ms sd

```

Private pages on a shared map —  _what am I, a wizard?_  On Linux, even a shared memory is counted as private unless it’s actually shared. Let’s see if it’s in the buffer cache:

```
# Seems like the first part is...
$ vmtouch -m 64G -v mordor.map
[OOo                                      ] 389440/8388608

           Files: 1
     Directories: 0
  Resident Pages: 389440/8388608  1G/32G  4.64%
         Elapsed: 0.27624 seconds

# Let's load it in the cache!
$ cat mordor.map > /dev/null
$ vmtouch -m 64G -v mordor.map
[ooooooooooooooooooooooooo      oooOOOOOOO] 2919606/8388608

           Files: 1
     Directories: 0
  Resident Pages: 2919606/8388608  11G/32G  34.8%
         Elapsed: 0.59845 seconds

```

Whoa, simply reading a file gets it cached? Anyway, how’s our process?

```
ps -p $$ -o pid,rss
  PID   RSS
17906 286584 # <-- Wait a bloody minute

```

A common misconception is that mapping a file consumes memory, whereas reading it using file API does not. One way or another, the pages from that file are going to get in the buffer cache. There is only a small difference, a process has to create the page table entries with mmap way, but the pages themselves are shared. Interestingly our process Rss shrinked, as there was  _a demand_  for the process pages.

#### Sometimes all of our thoughts are misgiven

The file-backed memory is always reclaimable, the only difference between dirty and clean — the dirty memory has to be cleaned before it can be reclaimed. So should you panick when a process consumes a lot of memory in  `top`? Start panicking (mildly) when a process has a lot of anonymous dirty pages — these can’t be reclaimed. If you see a very large growing anonymous mapping segment, you’re probably in trouble  _(and make it double)_. But the Rss or even Pss is not to be blindly trusted.

Another common mistake is to assume any relation between the process virtual memory and the memory consumption, or even treating all memory maps equally. Any reclaimable memory, is as good as a free one. To put it simply, it’s not going to fail your next memory allocation, but it  _may_  increase latency — let me explain.

The memory manager is making hard choices about what to keep in the physical memory, and what not. It may decide to page out a part of the process memory to swap in favour of more space for buffers, so the process has to page in  _that_  part on the next access. Fortunately it’s usually configurable. For example, there is an option called  [swappiness](https://en.wikipedia.org/wiki/Swappiness)  on Linux, that determines when should the kernel start paging out anonymous memory. A value of  `0`  means  _“until abso-fucking-lutely necessary”_.

## An end, once and for all

If you got here, I salute you! I started this article as a break from actual work, in a hope that simply explaining a  _thousand times_  explained concepts in a more accessible way is going to help me to organize thoughts, and help others in the process. It took me longer than expected. Way more.

I have nothing but utmost respect for writers, as it’s a tedious hair-pulling process of neverending edits and rewrites. Somewhere, Jeff Atwood has said that the best book about learning how to code is the one about building houses. I can’t remember where it was, so I can’t quote it. I could only add, that book about writing comes next. After all, that’s programming in it’s distilled form — writing stories, clear an concise.

_EDIT:_  I’ve fixed the stupid mistakes with  `alloca()`  and  `sizeof(char *) vs sizeof(char)`, thanks immibis and BonzaiThePenguin. Thanks sWvich for pointing out missing cast in  `slab + sizeof(struct slab)`. Obviously I should have run the article through static analysis, but I didn’t — lesson learned.

_Open question_  — is there anything better than Markdown code block, where I could show an annotated excerpt with a possibility to download the whole code block?

1.  A fancy way to say you’re going to get different data. [↩](https://marek.vavrusa.com/memory/#fnref:coherent)
    

Written on February 20, 2015


# Linux Memory Types  

[linux top命令VIRT,RES,SHR,DATA的含义](https://javawind.net/p131)

- VIRT：virtual memory usage 虚拟内存
  1. 进程“需要的”虚拟内存大小，包括进程使用的库、代码、数据等
  2. 假如进程申请100m的内存，但实际只使用了10m，那么它会增长100m，而不是实际的使用量

- RES：resident memory usage 常驻内存 [RSS – Resident Set Size 实际使用物理内存](https://cloud.tencent.com/developer/article/1407483)
  1. 进程当前使用的内存大小，但不包括swap out
  2. 包含其他进程的共享
  3. 如果申请100m的内存，实际使用10m，它只增长10m，与VIRT相反
  4. 关于库占用内存的情况，它只统计加载的库文件所占内存大小

- SHR：shared memory 共享内存
  1. 除了自身进程的共享内存，也包括其他进程的共享内存
  2. 虽然进程只使用了几个共享库的函数，但它包含了整个共享库的大小
  3. 计算某个进程所占的物理内存大小公式：RES – SHR
  4. swap out后，它将会降下来

- DATA
  1. 数据占用的内存。如果top没有显示，按f键可以显示出来。
  2. 真正的该程序要求的数据空间，是真正在运行中要使用的。


```sh
$ uname -a
Linux 5.4.0-77-generic #86~18.04.1-Ubuntu

# man top
   Linux Memory Types
       For  our purposes there are three types of memory, and one is optional.  First is physical memory, a limited resource where code and data must reside when executed or referenced.  Next is the optional swap file, where modified (dirty) memory can be saved
       and later retrieved if too many demands are made on physical memory.  Lastly we have virtual memory, a nearly unlimited resource serving the following goals:

          1. abstraction, free from physical memory addresses/limits
          2. isolation, every process in a separate address space
          3. sharing, a single mapping can serve multiple needs
          4. flexibility, assign a virtual address to a file

       Regardless of which of these forms memory may take, all are managed as pages (typically 4096 bytes) but expressed by default in top as KiB (kibibyte).  The memory discussed under topic `2c. MEMORY Usage' deals with physical memory and the swap  file  for
       the system as a whole.  The memory reviewed in topic `3. FIELDS / Columns Display' embraces all three memory types, but for individual processes.

       For each such process, every memory page is restricted to a single quadrant from the table below.  Both physical memory and virtual memory can include any of the four, while the swap file only includes #1 through #3.  The memory in quadrant #4, when mod‐
       ified, acts as its own dedicated swap file.

                                     Private | Shared
                                 1           |          2
            Anonymous  . stack               |
                       . malloc()            |
                       . brk()/sbrk()        | . POSIX shm*
                       . mmap(PRIVATE, ANON) | . mmap(SHARED, ANON)
                      -----------------------+----------------------
                       . mmap(PRIVATE, fd)   | . mmap(SHARED, fd)
          File-backed  . pgms/shared libs    |
                                 3           |          4

       The following may help in interpreting process level memory values displayed as scalable columns and discussed under topic `3a. DESCRIPTIONS of Fields'.

          %MEM - simply RES divided by total physical memory
          CODE - the `pgms' portion of quadrant 3
          DATA - the entire quadrant 1 portion of VIRT plus all
                 explicit mmap file-backed pages of quadrant 3
          RES  - anything occupying physical memory which, beginning with
                 Linux-4.5, is the sum of the following three fields:
                 RSan - quadrant 1 pages, which include any
                        former quadrant 3 pages if modified
                 RSfd - quadrant 3 and quadrant 4 pages
                 RSsh - quadrant 2 pages
          RSlk - subset of RES which cannot be swapped out (any quadrant)
          SHR  - subset of RES (excludes 1, includes all 2 & 4, some 3)
          SWAP - potentially any quadrant except 4
          USED - simply the sum of RES and SWAP
          VIRT - everything in-use and/or reserved (all quadrants)

       Note: Even though program images and shared libraries are considered private to a process, they will be accounted for as shared (SHR) by the kernel.

```

[^v2exmemory]: [踩坑记： go 服务内存暴涨](https://www.v2ex.com/t/666257?p=1)

[^vavrusamemory]: [What a C programmer should know about memory](https://marek.vavrusa.com/memory/)

[[译] Ｃ程序员该知道的内存知识 (1)](https://segmentfault.com/a/1190000022531638)

[[译] Ｃ程序员该知道的内存知识 (2)](https://segmentfault.com/a/1190000022545986)

[[译] Ｃ程序员该知道的内存知识 (3)](https://segmentfault.com/a/1190000022656717)

[[译] Ｃ程序员该知道的内存知识 (4)](https://segmentfault.com/a/1190000022721381)

[[译] 程序员应该知道的内存知识](https://jason2506.gitbooks.io/cpumemory/content/)

