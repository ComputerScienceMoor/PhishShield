package main

import "math"

func PredictXGB(input []float64) []float64 {
	var var0 float64
	if input[4] < 1.0 {
		var0 = -0.19998048
	} else {
		if input[6] < 3.0 {
			if input[2] < 1.0 {
				var0 = -0.19967428
			} else {
				if input[11] < 0.677 {
					var0 = 0.19090612
				} else {
					var0 = -0.08195763
				}
			}
		} else {
			if input[2] < 3.0 {
				if input[5] < 1.0 {
					var0 = -0.14098361
				} else {
					var0 = -0.19953734
				}
			} else {
				if input[6] < 4.0 {
					var0 = 0.19392559
				} else {
					var0 = -0.16786633
				}
			}
		}
	}
	var var1 float64
	if input[4] < 1.0 {
		var1 = -0.18185677
	} else {
		if input[6] < 3.0 {
			if input[2] < 1.0 {
				var1 = -0.18160054
			} else {
				if input[11] < 0.685 {
					var1 = 0.17315286
				} else {
					var1 = -0.10046541
				}
			}
		} else {
			if input[2] < 3.0 {
				if input[5] < 1.0 {
					var1 = -0.12754527
				} else {
					var1 = -0.18144682
				}
			} else {
				if input[6] < 5.0 {
					var1 = 0.15241651
				} else {
					var1 = -0.18065327
				}
			}
		}
	}
	var var2 float64
	if input[4] < 1.0 {
		var2 = -0.1682436
	} else {
		if input[6] < 3.0 {
			if input[2] < 1.0 {
				var2 = -0.16801484
			} else {
				if input[5] < 1.0 {
					var2 = 0.16281378
				} else {
					var2 = 0.020573484
				}
			}
		} else {
			if input[0] < 34.0 {
				if input[2] < 2.0 {
					var2 = -0.1440269
				} else {
					var2 = 0.14642352
				}
			} else {
				if input[6] < 4.0 {
					var2 = -0.13320911
				} else {
					var2 = -0.1673026
				}
			}
		}
	}
	var var3 float64
	if input[4] < 1.0 {
		var3 = -0.15767376
	} else {
		if input[6] < 3.0 {
			if input[2] < 1.0 {
				var3 = -0.15745926
			} else {
				if input[11] < 0.677 {
					var3 = 0.14967935
				} else {
					var3 = -0.08319842
				}
			}
		} else {
			if input[2] < 3.0 {
				if input[5] < 1.0 {
					var3 = -0.105896726
				} else {
					var3 = -0.15793942
				}
			} else {
				if input[6] < 5.0 {
					var3 = 0.1347594
				} else {
					var3 = -0.15579385
				}
			}
		}
	}
	var var4 float64
	if input[4] < 1.0 {
		var4 = -0.14925848
	} else {
		if input[6] < 3.0 {
			if input[2] < 1.0 {
				var4 = -0.14904973
			} else {
				if input[5] < 1.0 {
					var4 = 0.1437493
				} else {
					var4 = 0.0050835665
				}
			}
		} else {
			if input[0] < 34.0 {
				if input[2] < 2.0 {
					var4 = -0.1248441
				} else {
					var4 = 0.13427877
				}
			} else {
				if input[6] < 4.0 {
					var4 = -0.11453785
				} else {
					var4 = -0.14839095
				}
			}
		}
	}
	var var5 float64
	if input[4] < 1.0 {
		var5 = -0.14242618
	} else {
		if input[6] < 3.0 {
			if input[2] < 1.0 {
				var5 = -0.14221723
			} else {
				if input[5] < 2.0 {
					var5 = 0.13521935
				} else {
					var5 = -0.032252472
				}
			}
		} else {
			if input[0] < 34.0 {
				if input[2] < 2.0 {
					var5 = -0.11722565
				} else {
					var5 = 0.12304079
				}
			} else {
				if input[5] < 1.0 {
					var5 = -0.11062177
				} else {
					var5 = -0.14297356
				}
			}
		}
	}
	var var6 float64
	if input[4] < 1.0 {
		var6 = -0.13679162
	} else {
		if input[6] < 3.0 {
			if input[2] < 1.0 {
				var6 = -0.13657796
			} else {
				if input[11] < 0.642 {
					var6 = 0.12989186
				} else {
					var6 = -0.023598976
				}
			}
		} else {
			if input[5] < 1.0 {
				if input[11] < 0.606 {
					var6 = 0.09592303
				} else {
					var6 = -0.13642035
				}
			} else {
				if input[13] < 11.0 {
					var6 = -0.07460552
				} else {
					var6 = -0.1368664
				}
			}
		}
	}
	var var7 float64
	if input[4] < 1.0 {
		var7 = -0.13208516
	} else {
		if input[6] < 3.0 {
			if input[2] < 1.0 {
				var7 = -0.13186318
			} else {
				if input[5] < 1.0 {
					var7 = 0.12542227
				} else {
					var7 = -0.010221942
				}
			}
		} else {
			if input[6] < 4.0 {
				if input[2] < 2.0 {
					var7 = -0.104127966
				} else {
					var7 = 0.07503849
				}
			} else {
				if input[2] < 3.0 {
					var7 = -0.13132016
				} else {
					var7 = -0.10697426
				}
			}
		}
	}
	var var8 float64
	if input[4] < 1.0 {
		var8 = -0.12811226
	} else {
		if input[6] < 3.0 {
			if input[2] < 1.0 {
				var8 = -0.12787879
			} else {
				if input[11] < 0.677 {
					var8 = 0.11819889
				} else {
					var8 = -0.11006228
				}
			}
		} else {
			if input[5] < 1.0 {
				if input[11] < 0.586 {
					var8 = 0.10589784
				} else {
					var8 = -0.118070476
				}
			} else {
				if input[13] < 11.0 {
					var8 = -0.06569904
				} else {
					var8 = -0.12825006
				}
			}
		}
	}
	var var9 float64
	if input[4] < 1.0 {
		var9 = -0.1247288
	} else {
		if input[6] < 3.0 {
			if input[5] < 1.0 {
				if input[2] < 1.0 {
					var9 = -0.12433193
				} else {
					var9 = 0.11703034
				}
			} else {
				if input[6] < 2.0 {
					var9 = 0.124676585
				} else {
					var9 = -0.19243361
				}
			}
		} else {
			if input[5] < 1.0 {
				if input[11] < 0.606 {
					var9 = 0.08296518
				} else {
					var9 = -0.12062927
				}
			} else {
				if input[2] < 3.0 {
					var9 = -0.12539545
				} else {
					var9 = -0.08546352
				}
			}
		}
	}
	var var10 float64
	if input[4] < 1.0 {
		var10 = -0.12182571
	} else {
		if input[6] < 3.0 {
			if input[6] < 2.0 {
				if input[2] < 1.0 {
					var10 = -0.12133869
				} else {
					var10 = 0.12188061
				}
			} else {
				if input[11] < 0.613 {
					var10 = 0.076133065
				} else {
					var10 = -0.18201144
				}
			}
		} else {
			if input[5] < 1.0 {
				if input[11] < 0.56 {
					var10 = 0.112896286
				} else {
					var10 = -0.100057065
				}
			} else {
				if input[13] < 12.0 {
					var10 = -0.065195344
				} else {
					var10 = -0.12180562
				}
			}
		}
	}
	var var11 float64
	if input[4] < 1.0 {
		var11 = -0.1193188
	} else {
		if input[6] < 3.0 {
			if input[6] < 2.0 {
				if input[2] < 1.0 {
					var11 = -0.1186468
				} else {
					var11 = 0.11888552
				}
			} else {
				if input[2] < 2.0 {
					var11 = -0.05419317
				} else {
					var11 = 0.11983997
				}
			}
		} else {
			if input[6] < 4.0 {
				if input[2] < 2.0 {
					var11 = -0.08606293
				} else {
					var11 = 0.06967871
				}
			} else {
				if input[2] < 3.0 {
					var11 = -0.11760988
				} else {
					var11 = -0.093163915
				}
			}
		}
	}
	var var12 float64
	if input[4] < 1.0 {
		var12 = -0.117142096
	} else {
		if input[6] < 3.0 {
			if input[5] < 1.0 {
				if input[11] < 0.653 {
					var12 = 0.11001158
				} else {
					var12 = -0.054421127
				}
			} else {
				if input[6] < 2.0 {
					var12 = 0.1085096
				} else {
					var12 = -0.16883434
				}
			}
		} else {
			if input[5] < 1.0 {
				if input[11] < 0.606 {
					var12 = 0.07506438
				} else {
					var12 = -0.10974387
				}
			} else {
				if input[2] < 3.0 {
					var12 = -0.11770743
				} else {
					var12 = -0.070623055
				}
			}
		}
	}
	var var13 float64
	if input[4] < 1.0 {
		var13 = -0.11524304
	} else {
		if input[6] < 3.0 {
			if input[6] < 2.0 {
				if input[2] < 1.0 {
					var13 = -0.11769797
				} else {
					var13 = 0.114141546
				}
			} else {
				if input[2] < 2.0 {
					var13 = -0.051249564
				} else {
					var13 = 0.1152062
				}
			}
		} else {
			if input[5] < 1.0 {
				if input[11] < 0.606 {
					var13 = 0.06770576
				} else {
					var13 = -0.10622833
				}
			} else {
				if input[13] < 12.0 {
					var13 = -0.050083436
				} else {
					var13 = -0.11497756
				}
			}
		}
	}
	var var14 float64
	if input[4] < 1.0 {
		var14 = -0.11357935
	} else {
		if input[6] < 3.0 {
			if input[6] < 2.0 {
				if input[2] < 1.0 {
					var14 = -0.11545359
				} else {
					var14 = 0.11199911
				}
			} else {
				if input[11] < 0.565 {
					var14 = 0.07772275
				} else {
					var14 = -0.10808549
				}
			}
		} else {
			if input[6] < 4.0 {
				if input[2] < 2.0 {
					var14 = -0.07544405
				} else {
					var14 = 0.06591223
				}
			} else {
				if input[2] < 3.0 {
					var14 = -0.111027375
				} else {
					var14 = -0.08558701
				}
			}
		}
	}
	var var15 float64
	if input[4] < 1.0 {
		var15 = -0.112116516
	} else {
		if input[6] < 3.0 {
			if input[5] < 1.0 {
				if input[11] < 0.642 {
					var15 = 0.10373439
				} else {
					var15 = -0.03771126
				}
			} else {
				if input[6] < 2.0 {
					var15 = 0.095941715
				} else {
					var15 = -0.14946786
				}
			}
		} else {
			if input[5] < 1.0 {
				if input[11] < 0.606 {
					var15 = 0.06451606
				} else {
					var15 = -0.10061874
				}
			} else {
				if input[2] < 3.0 {
					var15 = -0.11252053
				} else {
					var15 = -0.056972533
				}
			}
		}
	}
	var var16 float64
	if input[4] < 1.0 {
		var16 = -0.11082616
	} else {
		if input[6] < 3.0 {
			if input[6] < 2.0 {
				if input[2] < 1.0 {
					var16 = -0.114795744
				} else {
					var16 = 0.10852953
				}
			} else {
				if input[2] < 2.0 {
					var16 = -0.05007718
				} else {
					var16 = 0.110811666
				}
			}
		} else {
			if input[5] < 1.0 {
				if input[11] < 0.56 {
					var16 = 0.09112474
				} else {
					var16 = -0.07939895
				}
			} else {
				if input[2] < 3.0 {
					var16 = -0.11099696
				} else {
					var16 = -0.05292443
				}
			}
		}
	}
	var var17 float64
	if input[4] < 1.0 {
		var17 = -0.10968468
	} else {
		if input[6] < 2.0 {
			if input[2] < 1.0 {
				var17 = -0.1128816
			} else {
				if input[5] < 5.0 {
					var17 = 0.10724918
				} else {
					var17 = -0.36758062
				}
			}
		} else {
			if input[11] < 0.565 {
				if input[12] < 0.036 {
					var17 = 0.09100275
				} else {
					var17 = -0.100532435
				}
			} else {
				if input[1] < 26.0 {
					var17 = -0.14421768
				} else {
					var17 = -0.05902618
				}
			}
		}
	}
	var var18 float64
	if input[4] < 1.0 {
		var18 = -0.108672336
	} else {
		if input[6] < 2.0 {
			if input[2] < 1.0 {
				var18 = -0.111169696
			} else {
				if input[5] < 5.0 {
					var18 = 0.105723895
				} else {
					var18 = -0.27318287
				}
			}
		} else {
			if input[5] < 1.0 {
				if input[11] < 0.613 {
					var18 = 0.07586241
				} else {
					var18 = -0.10463815
				}
			} else {
				if input[13] < 11.0 {
					var18 = 0.00833252
				} else {
					var18 = -0.11580073
				}
			}
		}
	}
	var var19 float64
	if input[4] < 1.0 {
		var19 = -0.10777245
	} else {
		if input[6] < 2.0 {
			if input[2] < 1.0 {
				var19 = -0.10962819
			} else {
				if input[5] < 5.0 {
					var19 = 0.10428983
				} else {
					var19 = -0.21864402
				}
			}
		} else {
			if input[5] < 1.0 {
				if input[11] < 0.613 {
					var19 = 0.071154214
				} else {
					var19 = -0.09940211
				}
			} else {
				if input[13] < 11.0 {
					var19 = 0.0074936696
				} else {
					var19 = -0.11345821
				}
			}
		}
	}
	var var20 float64
	if input[4] < 1.0 {
		var20 = -0.10697087
	} else {
		if input[6] < 2.0 {
			if input[2] < 1.0 {
				var20 = -0.108230405
			} else {
				if input[12] < 0.18 {
					var20 = 0.103104174
				} else {
					var20 = -0.12307494
				}
			}
		} else {
			if input[5] < 1.0 {
				if input[11] < 0.642 {
					var20 = 0.060669895
				} else {
					var20 = -0.110895954
				}
			} else {
				if input[13] < 11.0 {
					var20 = 0.0067403503
				} else {
					var20 = -0.11114708
				}
			}
		}
	}
	var var21 float64
	if input[4] < 1.0 {
		var21 = -0.106255494
	} else {
		if input[6] < 2.0 {
			if input[2] < 1.0 {
				var21 = -0.106953695
			} else {
				if input[12] < 0.18 {
					var21 = 0.101811424
				} else {
					var21 = -0.10437135
				}
			}
		} else {
			if input[2] < 2.0 {
				if input[11] < 0.536 {
					var21 = -0.008459669
				} else {
					var21 = -0.09927802
				}
			} else {
				if input[6] < 4.0 {
					var21 = 0.096372016
				} else {
					var21 = -0.09130438
				}
			}
		}
	}
	var var22 float64
	if input[4] < 1.0 {
		var22 = -0.105615936
	} else {
		if input[6] < 2.0 {
			if input[2] < 1.0 {
				var22 = -0.10577869
			} else {
				if input[12] < 0.18 {
					var22 = 0.10055933
				} else {
					var22 = -0.08967291
				}
			}
		} else {
			if input[5] < 1.0 {
				if input[11] < 0.613 {
					var22 = 0.06254316
				} else {
					var22 = -0.08841476
				}
			} else {
				if input[13] < 12.0 {
					var22 = -0.010252093
				} else {
					var22 = -0.109528676
				}
			}
		}
	}
	var var23 float64
	if input[4] < 1.0 {
		var23 = -0.105043225
	} else {
		if input[6] < 2.0 {
			if input[2] < 1.0 {
				var23 = -0.1046886
			} else {
				if input[5] < 5.0 {
					var23 = 0.09912882
				} else {
					var23 = -0.14388159
				}
			}
		} else {
			if input[5] < 1.0 {
				if input[11] < 0.642 {
					var23 = 0.053087663
				} else {
					var23 = -0.10152471
				}
			} else {
				if input[13] < 12.0 {
					var23 = -0.0092535205
				} else {
					var23 = -0.10746296
				}
			}
		}
	}
	var var24 float64
	if input[4] < 1.0 {
		var24 = -0.10452955
	} else {
		if input[6] < 2.0 {
			if input[2] < 1.0 {
				var24 = -0.1036688
			} else {
				if input[12] < 0.18 {
					var24 = 0.0981266
				} else {
					var24 = -0.07716472
				}
			}
		} else {
			if input[2] < 2.0 {
				if input[11] < 0.536 {
					var24 = -0.008390145
				} else {
					var24 = -0.09188074
				}
			} else {
				if input[6] < 4.0 {
					var24 = 0.093229786
				} else {
					var24 = -0.08646025
				}
			}
		}
	}
	var var25 float64
	var25 = sigmoid(var0 + var1 + var2 + var3 + var4 + var5 + var6 + var7 + var8 + var9 + var10 + var11 + var12 + var13 + var14 + var15 + var16 + var17 + var18 + var19 + var20 + var21 + var22 + var23 + var24)
	return []float64{1.0 - var25, var25}
}
func sigmoid(x float64) float64 {
	if x < 0.0 {
		z := math.Exp(x)
		return z / (1.0 + z)
	}
	return 1.0 / (1.0 + math.Exp(-x))
}
