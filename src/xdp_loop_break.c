#define KBUILD_MODNAME "foo"
#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/ipv6.h>
#include <linux/types.h>

BPF_TABLE("array", int, long, dropcnt, 256);

int xdp_prog(struct xdp_md *xdp) {
  void *data_end = (void *)(long)xdp->data_end;
  void *data = (void *)(long)xdp->data;
  struct ethhdr *eth = data;
  __u16 h_proto;
  int i;
  long *value;

  if (data + sizeof(struct ethhdr) > data_end)
    return XDP_DROP;

  h_proto = eth->h_proto;
  int key = h_proto;

#pragma unroll
  for (i = 0; i < 10; ++i) {
    value = dropcnt.lookup(&key);
    if (value)
      break;
  };

  if (value)
    lock_xadd(value, 1);

  return XDP_DROP;
}
