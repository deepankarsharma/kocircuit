# FUNCTIONS WITH INTERMEDIATE STEPS

So far we've defined functions that simply invoke another function
and return its result. For example, we looked at:

	CustomFormalGreeting(firstName, familyName) {
		return: strings.Join(
			string: ("Hello", firstName, familyName)
			delimiter: " "
		)
	}

Suppose now that we want to compute some intermediate values,
before we compute the final return value.

For instance, suppose we want to write a function that produces
two greetings, e.g. "Hello John and Hello Mary". 
We want to use `CustomFormalGreeting` twice, to prepare
the intermediate values "Hello John" and "Hello Mary".
And then we would like to join these into the final greeting.
The following code does that:

	DoubleGreeting(name1, name2) {
		firstGreeting: CustomFormalGreeting(firstName: name1)
		secondGreeting: CustomFormalGreeting(firstName: name2)
		return: strings.Join(
			string: (firstGreeting, "and", secondGreeting)
			delimiter: " "
		)
	}

`firstGreeting` and `secondGreeting` are purely names for the 
values computed by calling `CustomFormalGreeting(firstName: name1)`
and `CustomFormalGreeting(firstName: name2)`, respectivelly.
In other words, they can be viewed as "immutable" or "assign-once" variables.

The Ko interpreter will recognize that the values named
`firstGreeting` and `secondGreeting` are needed by the `return:` expression
and it will ensure they are computed before the invocation to `strings.Join`
can take place.

In general, one can define any number of intermediate values, which we call _steps_,
by adding an expression of the following form, on a separate line in the function body (between
the curly braces):

	<label>: <formula>

Step `<label>`s need to be unique and different from the 
label `return`, reserved for the function's return value.

## EXERCISE

To try `DoubleGreeting`, add the function `GreetJohnAndMary` and execute it:

	GreetJohnAndMary() {
		return: DoubleGreeting(name1: "John", name2: "Mary")
	}

You will get the output:

	"Hello John and Hello Mary"