def remove_flags(args, flags):
    flag_values = flags.values()
    return filter(lambda arg: arg not in flag_values, args)



def check_value(value, type_def):
    return str(type(value)) == "<type '{}'>".format(type_def)
